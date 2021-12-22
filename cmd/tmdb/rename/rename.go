package rename

import (
	"fmt"
	"os"
	"io"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
	"github.com/ssuareza/tmdb"
	. "github.com/ssuareza/tmdb/cmd/tmdb/config"
)

func Command() *cobra.Command {
	var move bool

	var cmd = &cobra.Command{
		Use:   "rename <movie-file>",
		Short: "Rename movie file",
		Long: `Rename movie file based on TheMovieDB database.

Example:
  tmdb rename Joker.2019.720p.BluRay.x264-[YTS.LT].avi --move /media/Movies
  File renamed to "Joker (2019).avi"
  File moved to "/media/Movies/Joker (2019)/Joker (2019).avi"`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(os.Args) < 3 {
				return fmt.Errorf("%s", "you should pass a file")
			}
			if _, err := os.Stat(os.Args[2]); err != nil {
				return fmt.Errorf("file \"%s\" does not exist", os.Args[2])
			}

			if move {
				if len(os.Args) != 5 {
					return fmt.Errorf("with \"--move\" you should define a destination path")
				}
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			file := os.Args[2]

			results, err := searchName(cleanFilename(filepath.Base(file)))
			if len(results) == 0 {
				fmt.Println("No matches found")
				os.Exit(0)
			}
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}

			// rename
			newName := fmt.Sprintf("%s (%s)%s", results[0].Title, strings.Split(results[0].ReleaseDate, "-")[0], filepath.Ext(file))
			if err := os.Rename(file, filepath.Dir(file)+"/"+newName); err != nil {
				fmt.Printf("Not possible to rename file to \"%s\"", newName)
				os.Exit(0)
			}
			fmt.Printf("File renamed to \"%s\"\n", newName)

			// and move
			if move {
				dstDir := filepath.Clean(os.Args[4]) + "/" + newName[0:len(newName)-len(filepath.Ext(file))]
				dst := dstDir + "/" + newName
				if err := moveFile(filepath.Dir(file)+"/"+newName, dstDir); err != nil {
					fmt.Println(err)
					os.Exit(0)
				}

				fmt.Printf("File moved to \"%s\"\n", dst)
			}
		},
	}

	cmd.Flags().BoolVarP(&move, "move", "m", false, "Move file to another destination")
	return cmd
}

func cleanFilename(file string) string {
	extension := filepath.Ext(file)
	name := file[0 : len(file)-len(extension)]

	clean1 := strings.ReplaceAll(name, ".", " ")
	clean2 := strings.ReplaceAll(clean1, "-", " ")
	clean3 := strings.ReplaceAll(clean2, "(", " ")
	clean4 := strings.ReplaceAll(clean3, ")", " ")
	clean5 := strings.ReplaceAll(clean4, "[", " ")
	clean6 := strings.ReplaceAll(clean5, "]", " ")
	clean7 := strings.ReplaceAll(clean6, "  ", " ")
	return strings.TrimSpace(clean7)
}

func searchName(str string) ([]themoviedb.Movie, error) {
	// getting year if exists
	r, _ := regexp.Compile("[1-2][0-9][0-9][0-9]")
	year := r.FindString(str)

	var found bool
	var words []string
	var query string
	for found || len(str) != 0 {
		query = strings.Replace(strings.TrimSpace(str), " ", "%20", -1)
		client := themoviedb.NewClient(fmt.Sprint(APIKey))
		movies, err := client.SearchMovie(query, year)
		if err != nil {
			return []themoviedb.Movie{}, fmt.Errorf("not possible to perform a search")
		}
		if len(movies) != 0 {
			return movies, nil
		}

		// prepare next iteration
		words = strings.Fields(str)
		if len(words) == 0 {
			str = ""
		} else {
			str = strings.Replace(str, words[len(words)-1], "", -1)
		}
	}

	return []themoviedb.Movie{}, nil
}

func moveFile(file string, dst string) error {
	if err := os.MkdirAll(dst, os.ModePerm); err != nil {
		return fmt.Errorf("not possible to create \"%s\"", dst)
	}

	/*
   Move files with os.Rename() give error "invalid cross-device link" for Docker container with Volumes.
   We fix this creating a temporary (os.Create) file and making a copy (io.Copy).
	*/
	inputFile, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("Couldn't open source file: %s", err)
	}
	outputFile, err := os.Create(dst+"/"+filepath.Base(file))
	if err != nil {
		inputFile.Close()
		return fmt.Errorf("Couldn't open dest file: %s", err)
	}
	defer outputFile.Close()
	_, err = io.Copy(outputFile, inputFile)
	inputFile.Close()
	if err != nil {
		return fmt.Errorf("Writing to output file failed: %s", err)
	}
	// The copy was successful, so now delete the original file
	err = os.Remove(file)
	if err != nil {
		return fmt.Errorf("Failed removing original file: %s", err)
	}

	return nil
}

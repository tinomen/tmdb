package rename

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
	"github.com/ssuareza/themoviedb-cli"
	. "github.com/ssuareza/themoviedb-cli/cmd/tmdb/config"
)

func Command() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "rename",
		Short: "Rename",
		Long:  `Rename`,
		Run: func(cmd *cobra.Command, args []string) {
			str := cleanFilename(os.Args[2])
			results := searchName(str)
			fmt.Println(results[0])
		},
	}

	return cmd
}

func cleanFilename(filename string) string {
	extension := filepath.Ext(filename)
	name := filename[0 : len(filename)-len(extension)]

	clean1 := strings.ReplaceAll(name, ".", " ")
	clean2 := strings.ReplaceAll(clean1, "-", " ")
	clean3 := strings.ReplaceAll(clean2, "(", " ")
	clean4 := strings.ReplaceAll(clean3, ")", " ")
	clean5 := strings.ReplaceAll(clean4, "[", " ")
	clean6 := strings.ReplaceAll(clean5, "]", " ")
	clean7 := strings.ReplaceAll(clean6, "  ", " ")
	return clean7
}

func searchName(str string) []themoviedb.Movie {
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
			log.Fatal(err)
		}
		if len(movies) != 0 {
			return movies
		}

		// prepare next iteration
		words = strings.Fields(str)
		if len(words) == 0 {
			str = ""
		} else {
			str = strings.Replace(str, words[len(words)-1], "", -1)
		}
	}

	return []themoviedb.Movie{}
}

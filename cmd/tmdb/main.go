package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/ssuareza/themoviedb-cli"
	"github.com/ssuareza/themoviedb-cli/cmd/tmdb/rename"
	"github.com/ssuareza/themoviedb-cli/cmd/tmdb/search"
	"log"
)

func main() {
	cmd := &cobra.Command{
		Use:   "tmdb",
		Short: "The Movie Database CLI",
		Long:  "Description here...",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Do stuff here...")
			client := themoviedb.NewClient("236f35f8f18ba49fe7a6369e8733ef2e")
			movies, err := client.SearchMovie("Avengers")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(movies)

		},
	}

	cmd.AddCommand(search.Command())
	cmd.AddCommand(rename.Command())

	cmd.Execute()
}

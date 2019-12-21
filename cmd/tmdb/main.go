package main

import (
	"github.com/spf13/cobra"
	"github.com/ssuareza/themoviedb-cli/cmd/tmdb/rename"
	"github.com/ssuareza/themoviedb-cli/cmd/tmdb/search"
)

func main() {
	cmd := &cobra.Command{
		Use:   "tmdb",
		Short: "The Movie Database CLI",
		Long:  "Description here...",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	cmd.AddCommand(search.Command())
	cmd.AddCommand(rename.Command())

	cmd.Execute()
}

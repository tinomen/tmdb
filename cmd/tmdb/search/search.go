package search

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "search",
		Short: "Search",
		Long:  `Search`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Search")
		},
	}

	return cmd
}

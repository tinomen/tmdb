package rename

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "rename",
		Short: "Rename",
		Long:  `Rename`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Rename")
		},
	}

	return cmd
}

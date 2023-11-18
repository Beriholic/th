package cmd

import (
	"github.com/Beriholic/th/handler"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "ll", "l"},
	Short:   "list all files in trash",
	Run: func(cmd *cobra.Command, args []string) {
		handler.ShowTable()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

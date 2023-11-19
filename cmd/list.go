package cmd

import (
	"github.com/Beriholic/th/consts"
	"github.com/Beriholic/th/handler"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "ll", "l"},
	Short:   "list all files in trash",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return CheckFlag()
	},
	Run: func(cmd *cobra.Command, args []string) {
		handler.ShowTable(sortType, sortOrder)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringVarP(&sortType, "sort", "s", consts.DefaultSortType, "sort by name or date")
	listCmd.Flags().StringVarP(&sortOrder, "order", "o", consts.DefaultSortOrder, "sort order asc or desc")
}

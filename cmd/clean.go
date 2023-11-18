package cmd

import (
	"fmt"

	"github.com/Beriholic/trash/handler"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean all files in Trash",
	Run: func(cmd *cobra.Command, args []string) {
		err := handler.Clean()
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}

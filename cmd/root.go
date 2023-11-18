package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "trash",
	Short: "A CLI tool for managing trash files",
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("Done!")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

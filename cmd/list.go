/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"trash/handler"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all files in trash",
	Run: func(cmd *cobra.Command, args []string) {
		handler.ShowTable()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

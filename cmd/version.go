package cmd

import (
	"fmt"

	"github.com/Beriholic/th/consts"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("th %s\n", consts.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

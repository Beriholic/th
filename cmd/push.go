package cmd

import (
	"trash/handler"

	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push a file to trash",
	//获取多个参数
	Args: cobra.MinimumNArgs(1),
	// Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		for _, file := range args {
			err := handler.PushToTrash(file)
			if err != nil {
				return err
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)
}

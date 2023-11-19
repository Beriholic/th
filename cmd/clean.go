package cmd

import (
	"fmt"

	"github.com/Beriholic/th/consts/errs"
	"github.com/Beriholic/th/handler"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean all files in Trash",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("The trash will be cleaned, do you want to continue? [y/n]")
		var input string
		fmt.Scanln(&input)

		switch input {
		case "y", "Y":
			err := handler.Clean()
			if err != nil {
				fmt.Println(err)
			}
		case "n", "N":
			return
		default:
			fmt.Println(errs.ErrInvalidInput)
		}

	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}

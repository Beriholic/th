package cmd

import (
	"os"

	"github.com/Beriholic/th/consts"
	"github.com/Beriholic/th/consts/errs"
	"github.com/spf13/cobra"
)

var sortType string
var sortOrder string

var rootCmd = &cobra.Command{
	Use:   "th",
	Short: "A CLI tool for managing trash files",
}

func init() {
	sortType = consts.DefaultSortType
	sortOrder = consts.DefaultSortOrder
}
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func CheckFlag() error {
	if sortType != "name" && sortType != "date" {
		return errs.ErrInvaildSortType
	}
	if sortOrder != "asc" && sortOrder != "desc" {
		return errs.ErrInvalidSortOrder
	}
	return nil
}

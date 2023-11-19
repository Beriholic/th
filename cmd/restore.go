package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Beriholic/th/handler"

	"github.com/spf13/cobra"
)

var fileInfos []handler.Info
var isEmpty bool = false

var restoreCmd = &cobra.Command{
	Use:     "restore",
	Aliases: []string{"r", "rs", "res"},
	Short:   "restore a file from trash",
	Args:    cobra.MinimumNArgs(0),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if err := CheckFlag(); err != nil {
			return err
		}

		if len(args) == 0 {
			PreHandleRestoreWithNoArgs(sortType, sortOrder)
		} else {
			PreHandleRestoreWithArgs()
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if isEmpty {
			return
		}
		if len(args) == 0 {
			HandleRestoreWithNoArgs()
		} else {
			fmt.Println("ok")
			nums := make([]int, 0)
			for _, v := range args {
				num, err := strconv.Atoi(v)
				if err != nil {
					fmt.Println(err)
					return
				}
				nums = append(nums, num)
			}
			HandleRestoreWithArgs(nums)
		}
	},
}

func PreHandleRestoreWithArgs() {
	var err error
	fileInfos, err = handler.GetTrashList(sortOrder, sortType)
	if err != nil {
		fmt.Println(err)
	}
	if len(fileInfos) == 0 {
		isEmpty = true
	}
}
func PreHandleRestoreWithNoArgs(sortType, sortOrder string) {
	var err error
	fileInfos, err = handler.ShowTable(sortType, sortOrder)
	if err != nil {
		fmt.Println(err)
	}

	if len(fileInfos) == 0 {
		isEmpty = true
	}

}
func HandleRestoreWithArgs(num []int) {
	if err := handler.Restore(num, fileInfos); err != nil {
		fmt.Println(err)
		return
	}
}
func HandleRestoreWithNoArgs() {
	fmt.Println("input the id of the file you want to restore")
	nums := make([]int, 0)
	input := Read()
	for _, v := range strings.Split(input, " ") {
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
			return
		}
		nums = append(nums, num)
	}

	if err := handler.Restore(nums, fileInfos); err != nil {
		fmt.Println(err)
		return
	}

}

func Read() string {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.Trim(str, "\n")
	str = strings.Trim(str, " ")
	return str
}

func init() {
	rootCmd.AddCommand(restoreCmd)
	restoreCmd.Flags().StringVarP(&sortType, "sort", "s", "name", "sort by name or date")
	restoreCmd.Flags().StringVarP(&sortOrder, "order", "o", "asc", "sort order asc or desc")
}

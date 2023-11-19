package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Beriholic/th/handler"
	"github.com/spf13/cobra"
)

var delCmd = &cobra.Command{
	Use:        "del",
	ArgAliases: []string{"d"},
	Short:      "del files in trash",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if err := CheckFlag(); err != nil {
			return err
		}

		if len(args) == 0 {
			PreHandleDelWithNoArgs(sortType, sortOrder)
		} else {
			PreHandleDelWithArgs()
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if isEmpty {
			return
		}
		if len(args) == 0 {
			HandleDelWithNoArgs()
		} else {
			nums := make([]int, 0)
			for _, v := range args {
				num, err := strconv.Atoi(v)
				if err != nil {
					fmt.Println(err)
					return
				}
				nums = append(nums, num)
			}
			HandleDelWithArgs(nums)
		}

	},
}

func init() {
	rootCmd.AddCommand(delCmd)
	delCmd.Flags().StringVarP(&sortType, "sort", "s", "name", "sort by name or date")
	delCmd.Flags().StringVarP(&sortOrder, "order", "o", "asc", "sort order asc or desc")
}

func PreHandleDelWithArgs() {
	var err error
	fileInfos, err = handler.GetTrashList(sortType, sortOrder)
	if err != nil {
		fmt.Println(err)
	}
	if len(fileInfos) == 0 {
		isEmpty = true
	}
}
func PreHandleDelWithNoArgs(sortType, sortOrder string) {
	var err error
	fileInfos, err = handler.ShowTable(sortType, sortOrder)
	if err != nil {
		fmt.Println(err)
	}

	if len(fileInfos) == 0 {
		isEmpty = true
	}

}
func HandleDelWithArgs(num []int) {
	if err := handler.Delete(num, fileInfos); err != nil {
		fmt.Println(err)
		return
	}
}
func HandleDelWithNoArgs() {
	fmt.Println("input the id of the file you want to delete")
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

	if err := handler.Delete(nums, fileInfos); err != nil {
		fmt.Println(err)
		return
	}
}

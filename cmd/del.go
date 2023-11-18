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
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			PostHandleDelWithNoArgs()
		} else {
			PostHandleDelWithArgs()
		}
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
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("Done!")
	},
}

func init() {
	rootCmd.AddCommand(delCmd)
}

func PostHandleDelWithArgs() {
	var err error
	fileInfos, err = handler.GetTrashList()
	if err != nil {
		fmt.Println(err)
	}
	if len(fileInfos) == 0 {
		isEmpty = true
	}
}
func PostHandleDelWithNoArgs() {
	var err error
	fileInfos, err = handler.ShowTable()
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
	input = strings.Replace(input, "\n", "", -1)
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

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
	Aliases: []string{"r", "rs"},
	Short:   "restore a file from trash",
	Args:    cobra.MinimumNArgs(0),
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			PostHandleRestoreWithNoArgs()
		} else {
			PostHandleRestoreWithArgs()
		}
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
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("ALL Done!")
	},
}

func PostHandleRestoreWithArgs() {
	var err error
	fileInfos, err = handler.GetTrashList()
	if err != nil {
		fmt.Println(err)
	}
	if len(fileInfos) == 0 {
		isEmpty = true
	}
}
func PostHandleRestoreWithNoArgs() {
	var err error
	fileInfos, err = handler.ShowTable()
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
	input = strings.Replace(input, "\n", "", -1)
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
	return str
}

func init() {
	rootCmd.AddCommand(restoreCmd)
}

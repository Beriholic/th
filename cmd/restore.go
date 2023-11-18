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
			PostHandleWithNoArgs()
		} else {
			PostHandleWithArgs()
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		if isEmpty {
			return
		}
		if len(args) == 0 {
			HandleWithNoArgs()
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
			HandleWithArgs(nums)
		}
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("ALL Done!")
	},
}

func PostHandleWithArgs() {
	var err error
	fileInfos, err = handler.GetTrashList()
	if err != nil {
		fmt.Println(err)
	}
	if len(fileInfos) == 0 {
		isEmpty = true
	}
}
func PostHandleWithNoArgs() {
	var err error
	fileInfos, err = handler.ShowTable()
	if err != nil {
		fmt.Println(err)
	}

	if len(fileInfos) == 0 {
		isEmpty = true
	}

}
func HandleWithArgs(num []int) {
	if err := handler.Restore(num, fileInfos); err != nil {
		fmt.Println(err)
		return
	}
}
func HandleWithNoArgs() {
	fmt.Println("input the id of the file you want to restore")
	nums := make([]int, 0)
	input := Read()
	input = strings.Replace(input, "\n", "", -1)
	// fmt.Println("input --> ", input)
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

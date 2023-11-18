package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Beriholic/trash/handler"

	"github.com/spf13/cobra"
)

var fileInfos []handler.Info
var isEmpty bool = false

var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "restore a file from trash",
	PreRun: func(cmd *cobra.Command, args []string) {
		var err error
		fileInfos, err = handler.ShowTable()
		if err != nil {
			fmt.Println(err)
		}

		if len(fileInfos) == 0 {
			isEmpty = true
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		if isEmpty {
			return
		}
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
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("ALL Done!")
	},
}

func Read() string {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	return str
}

func init() {
	rootCmd.AddCommand(restoreCmd)
}

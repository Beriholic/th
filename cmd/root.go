package cmd

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

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

func Read() string {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.Trim(str, "\n")
	str = strings.Trim(str, " ")
	return str
}

func InputToNums(input string, infoLens int) ([]int, error) {
	nums := make([]int, 0)
	st := make(map[int]struct{})
	strs := strings.Split(input, " ")

	for _, str := range strs {
		if str[0] == '[' {
			var regix = regexp.MustCompile(`\[(\d+)-(\d+)\]`)
			for _, v := range regix.FindAllStringSubmatch(str, -1) {
				start, err := strconv.Atoi(v[1])
				if err != nil {
					return nil, errs.BuildInfo(errs.ErrInvalidInput, err)
				}
				end, err := strconv.Atoi(v[2])
				if err != nil {
					return nil, errs.BuildInfo(errs.ErrInvalidInput, err)
				}
				if start > end {
					end, start = start, end
				}
				if OutOfRange(end, infoLens) {
					return nil, errs.BuildInfo(errs.ErrArrayOutOfRange, fmt.Errorf("max index is %d | max length is %d", end, infoLens-1))
				}

				for i := start; i <= end; i++ {
					if _, ok := st[i]; !ok {
						nums = append(nums, i)
						st[i] = struct{}{}
					}
				}
			}

		} else {
			num, err := strconv.Atoi(str)
			if err != nil {
				return nil, errs.BuildInfo(errs.ErrInvalidInput, err)
			}
			if _, ok := st[num]; !ok {
				nums = append(nums, num)
				st[num] = struct{}{}
			}
		}
	}

	return nums, nil
}

func OutOfRange(maxIdx, maxLen int) bool {
	//maxLen is [a,b)
	return maxIdx >= maxLen
}

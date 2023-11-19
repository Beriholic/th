package handler

import (
	"fmt"

	"github.com/Beriholic/th/consts"
	"github.com/Beriholic/th/consts/errs"
)

func Delete(nums []int, infos []Info) error {
	fmt.Println("------")
	fmt.Println("the following files will be deleted:")
	for i := 0; i < min(len(nums), 5); i++ {
		fmt.Println(infos[nums[i]].fileName)
	}
	if len(nums) > 5 {
		fmt.Println("...")
	}
	fmt.Println("do you want to continue? [y/n]")
	var input string
	fmt.Scanln(&input)

	switch input {
	case "y", "Y":
		break
	case "n", "N":
		return nil
	default:
		return errs.ErrInvalidInput
	}
	fmt.Println("------")
	for _, v := range nums {
		if OutOfRange(v, 0, len(infos)-1) {
			fmt.Println("id -->", v, errs.ErrArrayOutOfRange)
			continue
		}

		trashPath := fmt.Sprintf("%s/%s", consts.TrashFile, infos[v].fileName)
		infoPath := fmt.Sprintf("%s/%s.TrashInfo", consts.TrashInfo, infos[v].fileName)

		if err := RemoveFile(trashPath); err != nil {
			return errs.BuildInfo(errs.ErrRemoveFile, err)
		}

		if err := RemoveFile(infoPath); err != nil {
			return errs.BuildInfo(errs.ErrRemoveFile, err)
		}
		fmt.Println("id --> ", v, "Done!")
	}

	return nil
}

func OutOfRange(idx, l, r int) bool {
	return idx < l || idx > r
}

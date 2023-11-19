package handler

import (
	"fmt"

	"github.com/Beriholic/th/consts"
	"github.com/Beriholic/th/consts/errs"
)

func Restore(nums []int, infos []Info) error {
	fmt.Println("------")
	fmt.Println("the following files will be restored:")
	for i := 0; i < min(len(nums), 5); i++ {
		fmt.Println(infos[nums[i]].fileName)
	}
	if len(nums) > 5 {
		fmt.Println(".........")
	}
	fmt.Printf("%v files will be restored", len(nums))
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
		fromPath := infos[v].fromPath
		trashPath := fmt.Sprintf("%s/%s", consts.TrashFile, infos[v].fileName)
		infoPath := fmt.Sprintf("%s/%s.TrashInfo", consts.TrashInfo, infos[v].fileName)

		if err := MoveFile(trashPath, fromPath); err != nil {
			return errs.BuildInfo(errs.ErrRestore, err)
		}

		if err := RemoveFile(infoPath); err != nil {
			return errs.BuildInfo(errs.ErrRemoveFile, err)
		}
		fmt.Println("id --> ", v, "Done!")
	}

	return nil
}

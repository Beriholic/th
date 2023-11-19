package handler

import (
	"fmt"

	"github.com/Beriholic/th/consts/errs"
)

func Restore(nums []int, infos []Info) error {
	fmt.Println("------")
	for _, v := range nums {
		if OutOfRange(v, 0, len(infos)-1) {
			fmt.Println("id --> ", v, "out of range")
			continue
		}

		fromPath := infos[v].fromPath
		trashPath := fmt.Sprintf("%s/%s", TrashFile, infos[v].fileName)
		infoPath := fmt.Sprintf("%s/%s.TrashInfo", TrashInfo, infos[v].fileName)

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

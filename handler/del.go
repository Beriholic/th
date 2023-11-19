package handler

import (
	"fmt"

	"github.com/Beriholic/th/consts/errs"
)

func Delete(nums []int, infos []Info) error {
	fmt.Println("------")
	for _, v := range nums {
		if OutOfRange(v, 0, len(infos)-1) {
			fmt.Println("id -->", v, errs.ErrArrayOutOfRange)
			continue
		}

		trashPath := fmt.Sprintf("%s/%s", TrashFile, infos[v].fileName)
		infoPath := fmt.Sprintf("%s/%s.TrashInfo", TrashInfo, infos[v].fileName)

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

package handler

import (
	"fmt"
)

func Restore(nums []int, infos []Info) error {
	fmt.Println("------")
	for _, v := range nums {
		fromPath := infos[v].fromPath
		trashPath := fmt.Sprintf("%s/%s", TrashFile, infos[v].fileName)
		infoPath := fmt.Sprintf("%s/%s.TrashInfo", TrashInfo, infos[v].fileName)

		if err := MoveFile(trashPath, fromPath); err != nil {
			return fmt.Errorf("还原文件失败 --> %v", err)
		}

		if err := RemoveFile(infoPath); err != nil {
			return fmt.Errorf("删除info文件失败 --> %v", err)
		}
		fmt.Println("id --> ", v, "Done!")
	}

	return nil
}

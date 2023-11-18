package handler

import "fmt"

func Delete(nums []int, infos []Info) error {
	fmt.Println("------")
	for _, v := range nums {
		if OutOfRange(v, 0, len(infos)-1) {
			fmt.Println("id --> ", v, "out of range")
			continue
		}

		trashPath := fmt.Sprintf("%s/%s", TrashFile, infos[v].fileName)
		infoPath := fmt.Sprintf("%s/%s.TrashInfo", TrashInfo, infos[v].fileName)

		if err := RemoveFile(trashPath); err != nil {
			return fmt.Errorf("删除文件失败 --> %v", err)
		}

		if err := RemoveFile(infoPath); err != nil {
			return fmt.Errorf("删除info文件失败 --> %v", err)
		}
		fmt.Println("id --> ", v, "Done!")
	}

	return nil
}

func OutOfRange(idx, l, r int) bool {
	return idx < l || idx > r
}

package handler

import (
	"fmt"
	"os"
)

func Clean() error {
	if err := cleanDir(TrashFile); err != nil {
		return err
	}
	if err := cleanDir(TrashInfo); err != nil {
		return err
	}
	return nil
}
func cleanDir(path string) error {
	dr, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("打开文件夹失败 --> %v", err)
	}
	defer dr.Close()

	files, err := dr.Readdir(0)
	if err != nil {
		return fmt.Errorf("读取文件夹失败 --> %v", err)
	}
	for _, name := range files {
		if err := os.RemoveAll(path + "/" + name.Name()); err != nil {
			return fmt.Errorf("删除文件失败 --> %v", err)
		}
	}

	return nil
}

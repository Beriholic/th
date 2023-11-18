package handler

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

var home = os.Getenv("HOME")

var TrashFile = fmt.Sprintf("%s/.local/share/Trash/files", home)
var Trashinfo = fmt.Sprintf("%s/.local/share/Trash/info", home)

func PushFileArrayToTrash(fileArray []string) error {
	for _, file := range fileArray {
		err := PushToTrash(file)
		if err != nil {
			return err
		}
	}
	return nil
}

func PushToTrash(path string) error {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return fmt.Errorf("获取文件绝对路径失败 --> %v", err)
	}
	//debug
	// fmt.Println(absPath)
	// fmt.Println(TrashFile + "/" + filepath.Base(absPath))
	//
	err = os.Rename(absPath, TrashFile+"/"+filepath.Base(absPath))
	if err != nil {
		return fmt.Errorf("移动文件失败 --> %v", err)
	}
	return SaveFileData(absPath)
}

func SaveFileData(path string) error {
	deletionData := time.Now().Format("2006-01-02T15:04:05")

	//在info目录下创建一个文件
	file, err := os.Create(Trashinfo + "/" + filepath.Base(path) + ".trashinfo")
	if err != nil {
		return fmt.Errorf("创建文件失败 --> %v", err)
	}
	file.WriteString("[Trash Info]\n")
	file.WriteString(fmt.Sprintf("Path=%s\n", path))
	file.WriteString(fmt.Sprintf("DeletionDate=%s\n", deletionData))
	file.Close()

	return nil
}

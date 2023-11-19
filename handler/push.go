package handler

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/Beriholic/th/consts/errs"
)

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
		return errs.BuildInfo(errs.ErrGetAbsPath, err)
	}
	err = MoveFile(absPath, TrashFile+"/"+filepath.Base(absPath))
	if err != nil {
		return errs.BuildInfo(errs.ErrMoveFile, err)
	}
	return SaveFileData(absPath)
}

func SaveFileData(path string) error {
	deletionData := time.Now().Format("2006-01-02T15:04:05")

	//在info目录下创建一个文件
	file, err := os.Create(TrashInfo + "/" + filepath.Base(path) + ".TrashInfo")
	if err != nil {
		return errs.BuildInfo(errs.ErrNewFile, err)
	}
	file.WriteString("[Trash Info]\n")
	file.WriteString(fmt.Sprintf("Path=%s\n", path))
	file.WriteString(fmt.Sprintf("DeletionDate=%s\n", deletionData))
	file.Close()

	return nil
}

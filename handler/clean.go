package handler

import (
	"os"

	"github.com/Beriholic/th/consts/errs"
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
		return errs.BuildInfo(errs.ErrOpenFile, err)
	}
	defer dr.Close()

	files, err := dr.Readdir(0)
	if err != nil {
		return errs.BuildInfo(errs.ErrReadFile, err)
	}
	for _, name := range files {
		if err := os.RemoveAll(path + "/" + name.Name()); err != nil {
			return errs.BuildInfo(errs.ErrRemoveFile, err)
		}
	}

	return nil
}

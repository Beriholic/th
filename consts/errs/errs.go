package errs

import (
	"errors"
	"fmt"
)

var (
	ErrOpenFile   = errors.New("open file failed")
	ErrReadFile   = errors.New("read file failed")
	ErrRemoveFile = errors.New("remove file failed")
	ErrMoveFile   = errors.New("move file failed")
	ErrMakeDir    = errors.New("make dir failed")
	ErrRestore    = errors.New("restore file failed")
	ErrNewFile    = errors.New("create new file failed")

	ErrArrayOutOfRange = errors.New("array out of range")

	ErrGetAbsPath = errors.New("get absolute path failed")
)

func BuildInfo(msg, err error) error {
	return fmt.Errorf("%s --> %v", msg, err)
}

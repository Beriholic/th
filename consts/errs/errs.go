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

	ErrInvaildSortType  = errors.New("sort type must be name or date")
	ErrInvalidSortOrder = errors.New("sort order must be asc or desc")
	ErrInvalidInput     = errors.New("invalid input")
)

func BuildInfo(msg, err error) error {
	return fmt.Errorf("%s --> %v", msg, err)
}

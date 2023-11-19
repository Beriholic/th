package consts

import (
	"fmt"
	"os"
)

var (
	DefaultSortType  = "name"
	DefaultSortOrder = "asc"
)

var (
	home      = os.Getenv("HOME")
	TrashFile = fmt.Sprintf("%s/.local/share/Trash/files", home)
	TrashInfo = fmt.Sprintf("%s/.local/share/Trash/info", home)
)

var (
	Version = "v1.0.1"
)

package handler

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"go.etcd.io/etcd/client/pkg/v3/fileutil"
	"gopkg.in/ini.v1"
)

var home = os.Getenv("HOME")

var TrashFile = fmt.Sprintf("%s/.local/share/Trash/files", home)
var TrashInfo = fmt.Sprintf("%s/.local/share/Trash/info", home)

type FileInfo struct {
	infos []Info
}

func NewFileInfo() *FileInfo {
	return &FileInfo{
		infos: []Info{},
	}
}

type Info struct {
	id        int
	fileName  string
	fromPath  string
	trashTime string
}

func GetTrashList() ([]Info, error) {
	fileInfo := NewFileInfo()
	files, err := fileutil.ReadDir(TrashInfo)
	if err != nil {
		return nil, fmt.Errorf("读取文件夹失败 --> %v", err)
	}
	id := 0

	for _, file := range files {
		path := fmt.Sprintf("%s/%s", TrashInfo, file)
		fileInfo.BuildInfo(id, path)
		id++
	}

	return fileInfo.infos, nil
}

func (f *FileInfo) BuildInfo(id int, path string) error {
	cfg, err := ini.Load(path)
	if err != nil {
		return fmt.Errorf("读取文件失败 --> %v", err)
	}
	fromPath := cfg.Section("Trash Info").Key("Path").String()
	_trashTime := cfg.Section("Trash Info").Key("DeletionDate").String()
	trashTime, _ := time.Parse("2006-01-02T15:04:05", _trashTime)

	f.infos = append(f.infos, Info{
		id:        id,
		fileName:  filepath.Base(fromPath),
		fromPath:  fromPath,
		trashTime: trashTime.Format("2006-01-02 15:04:05"),
	})

	// debug
	// fmt.Println("debug")
	// for _, v := range fileInfo {
	// 	fmt.Println(v)
	// }
	return nil
}

func MoveFile(from, to string) error {
	return os.Rename(from, to)
}

func RemoveFile(path string) error {
	return os.Remove(path)
}

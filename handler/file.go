package handler

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/Beriholic/th/consts"
	"github.com/Beriholic/th/consts/errs"
	"go.etcd.io/etcd/client/pkg/v3/fileutil"
	"gopkg.in/ini.v1"
)

type FileInfo struct {
	infos []Info
}

func NewFileInfo() *FileInfo {
	return &FileInfo{
		infos: []Info{},
	}
}

type Info struct {
	fileName  string
	fromPath  string
	trashTime string
}

func GetTrashList(sortType, sortOrder string) ([]Info, error) {
	fileInfo := NewFileInfo()
	files, err := fileutil.ReadDir(consts.TrashInfo)
	if err != nil {
		return nil, errs.BuildInfo(errs.ErrReadFile, err)
	}

	for _, file := range files {
		path := fmt.Sprintf("%s/%s", consts.TrashInfo, file)
		fileInfo.BuildInfo(path)
	}

	sort.Slice(fileInfo.infos, func(i, j int) bool {
		switch sortType {
		case "name":
			if sortOrder == "asc" {
				return fileInfo.infos[i].fileName < fileInfo.infos[j].fileName
			} else {
				return fileInfo.infos[i].fileName > fileInfo.infos[j].fileName
			}
		case "data":
			if sortOrder == "asc" {
				return fileInfo.infos[i].trashTime < fileInfo.infos[j].trashTime
			} else {
				return fileInfo.infos[i].trashTime > fileInfo.infos[j].trashTime
			}
		}
		return false
	})

	return fileInfo.infos, nil
}

func (f *FileInfo) BuildInfo(path string) error {
	cfg, err := ini.Load(path)
	if err != nil {
		return errs.BuildInfo(errs.ErrReadFile, err)
	}
	fromPath := cfg.Section("Trash Info").Key("Path").String()
	_trashTime := cfg.Section("Trash Info").Key("DeletionDate").String()
	trashTime, _ := time.Parse("2006-01-02T15:04:05", _trashTime)

	f.infos = append(f.infos, Info{
		fileName:  filepath.Base(fromPath),
		fromPath:  fromPath,
		trashTime: trashTime.Format("2006-01-02 15:04:05"),
	})

	return nil
}

func MoveFile(from, to string) error {
	return os.Rename(from, to)
}

func RemoveFile(path string) error {
	return os.Remove(path)
}

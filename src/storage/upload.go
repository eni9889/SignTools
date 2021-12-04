package storage

import (
	"SignTools/src/util"
	"encoding/json"
	"github.com/tus/tusd/pkg/handler"
	"io"
	"os"
	"sync"
	"time"
)

type Upload interface {
	GetId() string
	delete() error
	GetData() (ReadonlyFile, error)
	GetInfo() (handler.FileInfo, error)
	GetModTime() (time.Time, error)
}

type upload struct {
	id string
	mu sync.Mutex
	FileSystemBase
}

func (u *upload) GetId() string {
	return u.id
}

func (u *upload) GetModTime() (time.Time, error) {
	stat, err := u.Stat(FSName(u.id))
	if err != nil {
		return time.Time{}, err
	}
	return stat.ModTime(), nil
}

func (u *upload) GetData() (ReadonlyFile, error) {
	return u.GetFile(FSName(u.id))
}

func (u *upload) GetInfo() (handler.FileInfo, error) {
	file, err := u.GetFile(FSName(u.id + ".info"))
	info := handler.FileInfo{}
	if err != nil {
		return info, err
	}
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return info, err
	}
	if err := json.Unmarshal(fileBytes, &info); err != nil {
		return info, err
	}
	return info, nil
}

func newUpload(id string) *upload {
	return &upload{id: id, FileSystemBase: FileSystemBase{resolvePath: func(name FSName) string {
		return util.SafeJoinFilePaths(uploadsPath, string(name))
	}}}
}

func (u *upload) delete() error {
	u.mu.Lock()
	defer u.mu.Unlock()
	if err := os.RemoveAll(u.resolvePath(FSName(u.id))); err != nil {
		return err
	}
	if err := os.RemoveAll(u.resolvePath(FSName(u.id + ".info"))); err != nil {
		return err
	}
	return nil
}

func GetUploadsPath() string {
	return uploadsPath
}

package driver

import (
	"os"
	"time"
)

type QSObject struct {
	ObjectName  string
	ObjectSize  int64
	CreatedTime time.Time
	IsDirectory bool
}

func (o *QSObject) Name() string {
	return o.ObjectName
}

func (o *QSObject) Size() int64 {
	return o.ObjectSize
}

func (o *QSObject) Mode() os.FileMode {
	if o.IsDir() {
		return os.ModeDir
	}
	return os.ModePerm
}

func (o *QSObject) ModTime() time.Time {
	return o.CreatedTime
}

func (o *QSObject) IsDir() bool {
	return o.IsDirectory
}

func (o *QSObject) Sys() interface{} {
	return true
}

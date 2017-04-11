package driver

import (
	"os"
	"time"
)

// QSObject stores Object's meta data
type QSObject struct {
	ObjectName  string
	ObjectSize  int64
	CreatedTime time.Time
	IsDirectory bool
}

// Name returns object's name
func (o *QSObject) Name() string {
	return o.ObjectName
}

// Size returns object's size
func (o *QSObject) Size() int64 {
	return o.ObjectSize
}

// Mode returns object's mode
func (o *QSObject) Mode() os.FileMode {
	if o.IsDir() {
		return os.ModeDir
	}
	return os.ModePerm
}

// ModTime returns object's createdTime
func (o *QSObject) ModTime() time.Time {
	return o.CreatedTime
}

// IsDir returns whether the object is a directory
func (o *QSObject) IsDir() bool {
	return o.IsDirectory
}

// Sys underlying data source
func (o *QSObject) Sys() interface{} {
	return true
}

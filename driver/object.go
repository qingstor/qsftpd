// +-------------------------------------------------------------------------
// | Copyright (C) 2017 Yunify, Inc.
// +-------------------------------------------------------------------------
// | Licensed under the Apache License, Version 2.0 (the "License");
// | you may not use this work except in compliance with the License.
// | You may obtain a copy of the License in the LICENSE file, or at:
// |
// | http://www.apache.org/licenses/LICENSE-2.0
// |
// | Unless required by applicable law or agreed to in writing, software
// | distributed under the License is distributed on an "AS IS" BASIS,
// | WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// | See the License for the specific language governing permissions and
// | limitations under the License.
// +-------------------------------------------------------------------------

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

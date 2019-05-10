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
	"errors"
	"io/ioutil"
	"os"

	"github.com/yunify/qingstor-sdk-go/service"
	"github.com/yunify/qsftpd/context"
)

// QSUploadFile stores ObjectKey and it's TempFile
type QSUploadFile struct {
	ObjectKey string

	TempFile *os.File
}

// Write writes data into a temp file
func (f *QSUploadFile) Write(p []byte) (n int, err error) {
	return f.TempFile.Write(p)
}

// Read does nothing but return an error
func (f *QSUploadFile) Read(p []byte) (n int, err error) {
	return 0, errors.New(`Upload file not allowed to read`)
}

// Close temp file and put an object
func (f *QSUploadFile) Close() error {
	defer context.Logger.Debugf("Delete temp file: %s", f.TempFile.Name())
	defer os.Remove(f.TempFile.Name())
	defer f.TempFile.Close()

	context.Logger.Debugf("Upload file: %s", f.TempFile.Name())
	context.Logger.Debugf("Upload file for key: %s", f.ObjectKey)

	f.TempFile.Seek(0, 0)

	_, err := context.Bucket.PutObject(f.ObjectKey, &service.PutObjectInput{Body: f.TempFile})
	return err
}

// Seek file
func (f *QSUploadFile) Seek(offset int64, whence int) (int64, error) {
	return f.TempFile.Seek(offset, whence)
}

// NewQSUploadFile creates a QSUploadFile struct
func NewQSUploadFile(objectKey string) (*QSUploadFile, error) {
	file, err := ioutil.TempFile(context.Settings.CachePath, "qsftp-")
	if err != nil {
		return nil, err
	}
	context.Logger.Debugf("Created temp file: %s", file.Name())
	return &QSUploadFile{ObjectKey: objectKey, TempFile: file}, nil
}

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
	"io"

	"github.com/yunify/qsftpd/context"
)

// QSDownloadFile stores ObjectKey and it's Body
type QSDownloadFile struct {
	ObjectKey string

	Body io.ReadCloser
}

// Write does nothing but return an error
func (f *QSDownloadFile) Write(p []byte) (n int, err error) {
	return 0, errors.New(`Download file not allowed to write`)
}

// Read data from file body
func (f *QSDownloadFile) Read(p []byte) (n int, err error) {
	return f.Body.Read(p)
}

// Close file
func (f *QSDownloadFile) Close() error {
	return f.Body.Close()
}

// Seek file
func (f *QSDownloadFile) Seek(offset int64, whence int) (int64, error) {
	return 0, nil
}

// NewQSDownloadFile creates a QSDownloadFile struct
func NewQSDownloadFile(objectKey string) (*QSDownloadFile, error) {
	output, err := context.Bucket.GetObject(objectKey, nil)
	if err != nil {
		return nil, err
	}
	context.Logger.Debugf("Open object: %s", objectKey)
	return &QSDownloadFile{ObjectKey: objectKey, Body: output.Body}, nil
}

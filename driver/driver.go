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
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/pengsrc/go-shared/convert"
	"github.com/yunify/qingstor-sdk-go/request/errors"
	"github.com/yunify/qingstor-sdk-go/service"
	"github.com/yunify/qsftpd/client"
	"github.com/yunify/qsftpd/context"
)

// QSDriver stores prefix
type QSDriver struct {
	prefix string
}

// ChangeDirectory changes current directory
func (d *QSDriver) ChangeDirectory(cc client.Context, directory string) error {
	directory = removeLeadingSlash(addTrailingSlash(trimPath(directory)))
	context.Logger.Debugf("Change directory: %s", directory)

	if directory == "" {
		return nil
	}

	d.prefix = fmt.Sprintf("/%s", directory)
	return nil
}

// MakeDirectory creates directory
func (d *QSDriver) MakeDirectory(cc client.Context, directory string) error {
	directory = removeLeadingSlash(addTrailingSlash(trimPath(directory)))
	context.Logger.Debugf("Mkdir directory: %s", directory)

	_, err := context.Bucket.PutObject(directory, &service.PutObjectInput{
		ContentType: convert.String("application/x-directory"),
	})
	return err
}

// ListFiles lists files in specified directory.
func (d *QSDriver) ListFiles(cc client.Context, dir string) ([]os.FileInfo, error) {
	if strings.HasSuffix(dir, "/-a") || strings.HasSuffix(dir, "/-l") {
		dir = dir[0 : len(dir)-2]
	}
	dir = trimPath(dir)
	if dir == "" {
		dir = cc.Path()
	}
	dir = removeLeadingSlash(addTrailingSlash(dir))

	context.Logger.Debugf("List files: %s", dir)

	infos := []os.FileInfo{}

	marker := convert.String("")
	for {
		output, err := context.Bucket.ListObjects(&service.ListObjectsInput{
			Prefix:    convert.String(dir),
			Delimiter: convert.String("/"),
			Marker:    marker,
		})
		if err != nil {
			return infos, err
		}

		for _, key := range output.Keys {
			if service.StringValue(key.Key) != dir {
				infos = append(infos, &QSObject{
					ObjectName:  trimPath(path.Base(service.StringValue(key.Key))),
					ObjectSize:  convert.Int64Value(key.Size),
					CreatedTime: convert.TimeValue(key.Created),
					IsDirectory: false,
				})
			}
		}

		for _, prefix := range output.CommonPrefixes {
			infos = append(infos, &QSObject{
				ObjectName:  trimPath(path.Base(service.StringValue(prefix))),
				ObjectSize:  int64(0),
				CreatedTime: time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC),
				IsDirectory: true,
			})
		}
		if convert.StringValue(output.NextMarker) == "" {
			break
		}
		marker = output.NextMarker
	}
	return infos, nil
}

// OpenFile opens file for read and write
func (d *QSDriver) OpenFile(cc client.Context, path string, flag int) (client.FileStream, error) {
	path = removeLeadingSlash(trimPath(path))
	context.Logger.Debugf("Open file: %s", path)

	switch flag {
	case os.O_RDONLY:
		return NewQSDownloadFile(path)
	case os.O_WRONLY:
		return NewQSUploadFile(path)
	}

	return nil, fmt.Errorf("Failed to open path: %s", path)
}

// DeleteFile delete a path
func (d *QSDriver) DeleteFile(cc client.Context, path string) error {
	path = removeLeadingSlash(trimPath(path))
	context.Logger.Debugf("Delete file: %s", path)

	_, err := context.Bucket.DeleteObject(path)
	if err != nil {
		return err
	}
	_, err = context.Bucket.DeleteObject(addTrailingSlash(path))
	if err != nil {
		return err
	}

	return nil
}

// GetFileInfo gets file stats
func (d *QSDriver) GetFileInfo(cc client.Context, filePath string) (os.FileInfo, error) {
	filePath = removeLeadingSlash(trimPath(filePath))
	context.Logger.Debugf("Get file info: %s", filePath)

	needTrailingSlash := false

	output, err := context.Bucket.HeadObject(filePath, nil)
	if err != nil {
		value, ok := err.(*errors.QingStorError)
		if !ok || value.StatusCode != 404 {
			return nil, err
		}
		needTrailingSlash = true
	}

	if needTrailingSlash {
		filePath = addTrailingSlash(filePath)
		output, err = context.Bucket.HeadObject(filePath, nil)
		if err != nil {
			value, ok := err.(*errors.QingStorError)
			if ok {
				if value.StatusCode == 404 {
					return nil, fmt.Errorf(`path "%s" not exists`, filePath)
				}
			}
			return nil, err
		}
	}

	return &QSObject{
		ObjectName:  trimPath(path.Base(filePath)),
		ObjectSize:  convert.Int64Value(output.ContentLength),
		CreatedTime: convert.TimeValue(output.LastModified),
		IsDirectory: needTrailingSlash,
	}, nil
}

// RenameFile renames a file name
func (d *QSDriver) RenameFile(cc client.Context, from, to string) error {
	from = trimPath(from)
	to = trimPath(to)
	if len(from) > 0 && strings.Index(from, "/") != 0 {
		from = d.prefix + from
	}
	if len(to) > 0 && strings.Index(to, "/") != 0 {
		to = d.prefix + to
	}
	from, err := filepath.Abs(from)
	if err != nil {
		return err
	}
	to, err = filepath.Abs(to)
	if err != nil {
		return err
	}
	if string(to[len(to)-1]) == "/" {
		to += path.Base(from)
	}
	from = removeLeadingSlash(from)
	to = removeLeadingSlash(to)
	context.Logger.Debugf("Rename file from: %s", from)
	context.Logger.Debugf("Rename file to: %s", to)

	needTrailingSlash := false

	_, err = context.Bucket.HeadObject(from, nil)
	if err != nil {
		value, ok := err.(*errors.QingStorError)
		if !ok || value.StatusCode != 404 {
			return err
		}
		needTrailingSlash = true
	}

	if needTrailingSlash {
		from = addTrailingSlash(from)
		to = addTrailingSlash(to)
		_, err = context.Bucket.HeadObject(from, nil)
		if err != nil {
			value, ok := err.(*errors.QingStorError)
			if !ok || value.StatusCode != 404 {
				return err
			}
		}
	}

	_, err = context.Bucket.PutObject(to, &service.PutObjectInput{
		XQSMoveSource: convert.String(fmt.Sprintf("/%s/%s", convert.StringValue(context.Bucket.Properties.BucketName), from)),
	})
	return err
}

// CanAllocate always returns true for the backend is QingStor Bucket
func (d *QSDriver) CanAllocate(cc client.Context, size int) (bool, error) {
	return true, nil
}

// ChmodFile changes file mode
func (d *QSDriver) ChmodFile(cc client.Context, path string, mode os.FileMode) error {
	return nil
}

func trimPath(path string) string {
	return strings.TrimSpace(strings.Replace(path, "/ ", "/", -1))
}

func removeLeadingSlash(path string) string {
	if len(path) > 0 {
		if strings.Index(path, "/") == 0 {
			return path[1:]
		}
		// Remove "D:\" and replace all "\" in filepath
		if strings.Index(path, "\\") == 2 {
			return strings.Replace(path[3:], "\\", "/", -1)
		}
	}
	return path
}

func addTrailingSlash(path string) string {
	if len(path) > 0 {
		if string(path[len(path)-1]) != "/" {
			return fmt.Sprintf(`%s/`, path)
		}
	}
	return path
}

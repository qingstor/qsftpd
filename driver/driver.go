package driver

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/pengsrc/go-utils/check"
	"github.com/pengsrc/go-utils/convert"
	"github.com/pengsrc/qsftp/client"
	"github.com/pengsrc/qsftp/context"
	"github.com/yunify/qingstor-sdk-go/request/errors"
	"github.com/yunify/qingstor-sdk-go/service"
)

type QSDriver struct {
	prefix string
}

func (d *QSDriver) ChangeDirectory(cc client.Context, directory string) error {
	directory = removeLeadingSlash(addTrailingSlash(trimPath(directory)))
	context.Logger.Debug("Change directory: %s", directory)

	if directory == "" {
		return nil
	}

	_, err := context.Bucket.HeadObject(directory, nil)
	if err != nil {
		value, ok := err.(*errors.QingStorError)
		if !ok && value.StatusCode != 404 {
			return err
		}

		output, err := context.Bucket.ListObjects(&service.ListObjectsInput{
			Prefix:    convert.String(directory),
			Delimiter: convert.String("/"),
		})
		if err != nil {
			return err
		}

		if len(output.Keys) == 0 && len(output.CommonPrefixes) == 0 {
			return fmt.Errorf(`directory "%s" not exists`, directory)
		}
	}

	d.prefix = fmt.Sprintf("/%s", directory)
	return nil
}

func (d *QSDriver) MakeDirectory(cc client.Context, directory string) error {
	directory = removeLeadingSlash(addTrailingSlash(trimPath(directory)))
	context.Logger.Debug("Mkdir directory: %s", directory)

	_, err := context.Bucket.PutObject(directory, &service.PutObjectInput{
		ContentType: convert.String("application/x-directory"),
	})
	return err
}

func (d *QSDriver) ListFiles(cc client.Context, dir string) ([]os.FileInfo, error) {
	dir = trimPath(dir)
	if dir == "" {
		dir = cc.Path()
	}
	dir = removeLeadingSlash(addTrailingSlash(dir))

	if check.StringSliceContains([]string{"-a", "-a/"}, dir) {
		dir = ""
	}

	context.Logger.Debug("List files: %s", dir)

	infos := []os.FileInfo{}

	output, err := context.Bucket.ListObjects(&service.ListObjectsInput{
		Prefix:    convert.String(dir),
		Delimiter: convert.String("/"),
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

	return infos, nil
}

func (d *QSDriver) OpenFile(cc client.Context, path string, flag int) (client.FileStream, error) {
	path = removeLeadingSlash(trimPath(path))
	context.Logger.Debug("Open file: %s", path)

	switch flag {
	case os.O_RDONLY:
		return NewQSDownloadFile(path)
	case os.O_WRONLY:
		return NewQSUploadFile(path)
	}

	return nil, fmt.Errorf("Failed to open path: %s", path)
}

func (d *QSDriver) DeleteFile(cc client.Context, path string) error {
	path = removeLeadingSlash(trimPath(path))
	context.Logger.Debug("Delete file: %s", path)

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

func (d *QSDriver) GetFileInfo(cc client.Context, filePath string) (os.FileInfo, error) {
	filePath = removeLeadingSlash(trimPath(filePath))
	context.Logger.Debug("Get file info: %s", filePath)

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
	context.Logger.Debug("Rename file from: %s", from)
	context.Logger.Debug("Rename file to: %s", to)

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

func (d *QSDriver) CanAllocate(cc client.Context, size int) (bool, error) {
	return true, nil
}

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

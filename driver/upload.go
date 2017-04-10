package driver

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/yunify/qsftp/context"
	"github.com/yunify/qingstor-sdk-go/service"
)

type QSUploadFile struct {
	ObjectKey string

	TempFile *os.File
}

func (f *QSUploadFile) Write(p []byte) (n int, err error) {
	return f.TempFile.Write(p)
}

func (f *QSUploadFile) Read(p []byte) (n int, err error) {
	return 0, errors.New(`Upload file not allowed to read`)
}

func (f *QSUploadFile) Close() error {
	defer context.Logger.Debug("Delete temp file: %s", f.TempFile.Name())
	defer os.Remove(f.TempFile.Name())
	defer f.TempFile.Close()

	context.Logger.Debug("Upload file: %s", f.TempFile.Name())
	context.Logger.Debug("Upload file for key: %s", f.ObjectKey)

	f.TempFile.Seek(0, 0)

	_, err := context.Bucket.PutObject(f.ObjectKey, &service.PutObjectInput{Body: f.TempFile})
	return err
}

func (f *QSUploadFile) Seek(offset int64, whence int) (int64, error) {
	return f.TempFile.Seek(offset, whence)
}

func NewQSUploadFile(objectKey string) (*QSUploadFile, error) {
	file, err := ioutil.TempFile("", "qsftp-")
	if err != nil {
		return nil, err
	}
	context.Logger.Debug("Created temp file: %s", file.Name())
	return &QSUploadFile{ObjectKey: objectKey, TempFile: file}, nil
}

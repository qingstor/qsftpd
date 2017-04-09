package driver

import (
	"errors"
	"io"

	"github.com/pengsrc/qsftp/context"
)

type QSDownloadFile struct {
	ObjectKey string

	Body io.ReadCloser
}

func (f *QSDownloadFile) Write(p []byte) (n int, err error) {
	return 0, errors.New(`Download file not allowed to write`)
}

func (f *QSDownloadFile) Read(p []byte) (n int, err error) {
	return f.Body.Read(p)
}

func (f *QSDownloadFile) Close() error {
	return f.Body.Close()
}

func (f *QSDownloadFile) Seek(offset int64, whence int) (int64, error) {
	return 0, nil
}

func NewQSDownloadFile(objectKey string) (*QSDownloadFile, error) {
	output, err := context.Bucket.GetObject(objectKey, nil)
	if err != nil {
		return nil, err
	}
	context.Logger.Debug("Open object: %s", objectKey)
	return &QSDownloadFile{ObjectKey: objectKey, Body: output.Body}, nil
}

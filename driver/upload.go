package driver

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/yunify/qingstor-sdk-go/service"
	"github.com/yunify/qsftp/context"
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
	defer context.Logger.DebugF("Delete temp file: %s", f.TempFile.Name())
	defer os.Remove(f.TempFile.Name())
	defer f.TempFile.Close()

	context.Logger.DebugF("Upload file: %s", f.TempFile.Name())
	context.Logger.DebugF("Upload file for key: %s", f.ObjectKey)

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
	file, err := ioutil.TempFile("", "qsftp-")
	if err != nil {
		return nil, err
	}
	context.Logger.DebugF("Created temp file: %s", file.Name())
	return &QSUploadFile{ObjectKey: objectKey, TempFile: file}, nil
}

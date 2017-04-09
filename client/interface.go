package client

import (
	"io"
	"os"
)

// Driver handles the file system access logic.
type Driver interface {
	// ChangeDirectory changes the current working directory.
	ChangeDirectory(cc Context, directory string) error

	// MakeDirectory creates a directory.
	MakeDirectory(cc Context, directory string) error

	// ListFiles lists the files of a directory.
	ListFiles(cc Context, dir string) ([]os.FileInfo, error)

	// OpenFile opens a file in 3 possible modes: read, write, appending write (use appropriate flags).
	OpenFile(cc Context, path string, flag int) (FileStream, error)

	// DeleteFile deletes a file or a directory.
	DeleteFile(cc Context, path string) error

	// GetFileInfo gets some info around a file or a directory.
	GetFileInfo(cc Context, path string) (os.FileInfo, error)

	// RenameFile renames a file or a directory.
	RenameFile(cc Context, from, to string) error

	// CanAllocate gives the approval to allocate some data.
	CanAllocate(cc Context, size int) (bool, error)

	// ChmodFile changes the attributes of the file.
	ChmodFile(cc Context, path string, mode os.FileMode) error
}

// Context is implemented on the server side to provide some access to
// few data around the client.
type Context interface {
	// Path provides the path of the current connection.
	Path() string
}

// FileStream is a read or write closeable stream.
type FileStream interface {
	io.Writer
	io.Reader
	io.Closer
	io.Seeker
}

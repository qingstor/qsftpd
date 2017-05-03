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

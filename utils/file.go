package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// RemoveContents will remove all contents under a folder.
func RemoveContents(dir string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, v := range files {
		err = os.RemoveAll(filepath.Join(dir, v.Name()))
		if err != nil {
			return err
		}
	}
	return nil
}

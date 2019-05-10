package utils

import (
	"github.com/magiconair/properties/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestRemoveContents(t *testing.T) {
	tempPath := os.TempDir() + "/qsftpd_test"
	err := os.MkdirAll(tempPath, 0775)
	if err != nil {
		t.Error(err)
		return
	}

	for i := 0; i < 100; i++ {
		f, err := ioutil.TempFile(tempPath, "")
		if err != nil {
			t.Error(err)
			return
		}
		f.Close()
	}

	files, err := ioutil.ReadDir(tempPath)
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, len(files), 100)

	err = RemoveContents(tempPath)
	if err != nil {
		t.Error(err)
		return
	}

	files, err = ioutil.ReadDir(tempPath)
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, len(files), 0)

	err = os.RemoveAll(tempPath)
	if err != nil {
		panic(err)
	}
}

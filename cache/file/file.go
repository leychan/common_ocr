package file

import (
	"errors"
	"io/ioutil"
	"os"

	gohelper "github.com/leychan/go-helper"
)

func Set(path string, body []byte) error {
	if !gohelper.FileOrDirExists(path) {
		gohelper.CreateFile(path)
	}
	err := gohelper.WriteFile(path, body)
	if err != nil {
		return err
	}
	return nil
}

func Get(path string) ([]byte, error) {
	if !gohelper.FileOrDirExists(path) {
		return []byte{}, errors.New("no token file")
	}
	f, err := os.Open(path)
	if err != nil {
		panic(err)
		return []byte{}, err
	}
	content, err := ioutil.ReadAll(f)
	if err != nil {
		return []byte{}, err
	}
	f.Close()
	return content, nil
}

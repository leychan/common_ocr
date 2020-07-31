package file

import (
	"errors"

	go_helper "github.com/leychan/go-helper"
)

func Set(path string, body []byte) (bool, error) {
	if !go_helper.FileOrDirExists(path) {
		go_helper.CreateFile(path)
	}
	err := go_helper.WriteFile(path, body, "")
	if err != nil {
		return false, err
	}
	return true, nil
}

func Get(path string) ([]byte, error) {
	if !go_helper.FileOrDirExists(path) {
		return []byte{}, errors.New("no token file")
	}

}

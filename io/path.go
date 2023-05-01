package io

import (
	"os"
	"path"
)

func Mkdir(filePath string) error {
	dir := path.Dir(filePath)
	if IsNotExist(dir) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func IsNotExist(path string) bool {
	_, err := os.Stat(path)
	return os.IsNotExist(err)
}

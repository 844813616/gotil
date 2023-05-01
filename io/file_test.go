package io

import (
	"log"
	"testing"
)

var testFilePath = "../temp/file_test"

func TestReadFile(t *testing.T) {
	ret, err := ReadFile(testFilePath)
	if err != nil {
		log.Println(err)
	}
	for _, v := range ret {
		println(v)
	}
}

func TestWriteFile(t *testing.T) {
	err := WriteFile(testFilePath, []string{"1", "2", "3", "4"})
	if err != nil {
		log.Println(err)
	}
}

func TestOpenFile(t *testing.T) {
	file, err := OpenFile(testFilePath, READ)
	if err != nil {
		return
	}
	err = file.Close()
	if err != nil {
		return
	}
}

package io

import (
	"log"
	"testing"
)

var testPath = "../temp/path_test"

func TestMkdir(t *testing.T) {
	err := Mkdir(testPath)
	if err != nil {
		log.Println(err)
	}
}

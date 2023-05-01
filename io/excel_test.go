package io

import (
	"log"
	"testing"
)

var testExcelPath = "../temp/excel_test.xlsx"

func TestWriteExcel(t *testing.T) {
	err := WriteExcel(testExcelPath, [][]string{{"1", "2", "3"}, {"2", "4", "6"}, {"3", "6", "9"}})
	if err != nil {
		log.Println(err)
	}
}

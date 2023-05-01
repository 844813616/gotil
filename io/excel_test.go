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

func TestReadExcel(t *testing.T) {
	content, err := ReadExcel(testExcelPath, defaultSheetName)
	if err != nil {
		log.Println(err)
	}
	for _, row := range content {
		for _, cell := range row {
			print(cell + " ")
		}
		println()
	}
}

func TestUpdateExcel(t *testing.T) {
	var m = make(map[string]bool)
	processor := func(in [][][]string) string {
		for _, a := range in {
			for _, r := range a {
				for _, c := range r {
					if m[c] == false {
						m[c] = true
						return c
					}
				}
			}
		}
		return ""
	}
	err := UpdateExcel(testExcelPath, defaultSheetName,
		Area{Coordinate{Absolute(1), Absolute(4)}, Coordinate{Absolute(3), Absolute(4)}},
		[]Area{
			{Coordinate{Absolute(1), Absolute(1)}, Coordinate{Absolute(3), Absolute(1)}},
			{Coordinate{Absolute(1), Absolute(2)}, Coordinate{Absolute(1), Absolute(3)}},
		}, processor)
	if err != nil {
		log.Println(err)
	}
}

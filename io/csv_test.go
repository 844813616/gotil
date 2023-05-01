package io

import (
	"log"
	"testing"
)

var testCsvPath = "../temp/csv_test.csv"

func TestReadCsv(t *testing.T) {
	records, err := ReadCsv(testCsvPath, false)
	if err != nil {
		log.Println(err)
		return
	}
	for _, record := range records {
		for _, v := range record {
			print(v + ",")
		}
		println()
	}
}

func TestWriteCsv(t *testing.T) {
	err := WriteCsv(testCsvPath, [][]string{{"1", "2", "3"}, {"2", "4", "6"}, {"3", "6", "9"}}, false)
	if err != nil {
		log.Println(err)
		return
	}
}

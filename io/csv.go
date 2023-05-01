package io

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
)

func ReadCsv(filePath string, skipHeader bool) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	bufReader := bufio.NewReader(file)
	csvReader := csv.NewReader(bufReader)
	if skipHeader {
		_, _ = csvReader.Read()
	}
	var records [][]string
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		records = append(records, record)
	}
	return records, nil
}

func WriteCsv(filePath string, records [][]string, append bool) error {
	var file *os.File
	var err error
	if append {
		file, err = os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	} else {
		file, err = os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	}
	if err != nil {
		return err
	}
	bufWriter := bufio.NewWriter(file)
	csvWriter := csv.NewWriter(bufWriter)
	if err = csvWriter.WriteAll(records); err != nil {
		return err
	}
	return nil
}

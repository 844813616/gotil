package io

import (
	"github.com/xuri/excelize/v2"
	"log"
)

var (
	defaultSheetName = "sheet1"
	defaultCell      = "A1"
)

func WriteExcel(filePath string, content [][]string) error {
	if err := Mkdir(filePath); err != nil {
		return err
	}
	file := excelize.NewFile()
	defer func() {
		if err := file.Close(); err != nil {
			log.Println(err)
		}
	}()
	_, err := file.NewSheet(defaultSheetName)
	if err != nil {
		return err
	}
	col, row, err := excelize.CellNameToCoordinates(defaultCell)
	if err != nil {
		return err
	}
	for i, line := range content {
		for j, value := range line {
			cellName, err := excelize.CoordinatesToCellName(col+j, row+i)
			if err != nil {
				log.Println(err)
				continue
			}
			if err := file.SetCellValue(defaultSheetName, cellName, value); err != nil {
				return err
			}
		}
	}
	if err := file.SaveAs(filePath); err != nil {
		return err
	}
	return nil
}

func UpdateExcel(filePath string, content map[string]string) error {
	return nil
}

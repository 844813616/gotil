package io

import (
	"encoding/xml"
	"github.com/xuri/excelize/v2"
	"log"
)

var (
	defaultSheetName           = "sheet1"
	defaultCell                = "A1"
	defaultXMLPathDocPropsCore = "docProps/core.xml"
	templateDocpropsCore       = `<cp:coreProperties xmlns:cp="http://schemas.openxmlformats.org/package/2006/metadata/core-properties" xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:dcterms="http://purl.org/dc/terms/" xmlns:dcmitype="http://purl.org/dc/dcmitype/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><dc:creator>weiguanqun</dc:creator><dcterms:created xsi:type="dcterms:W3CDTF">2006-09-16T00:00:00Z</dcterms:created><dcterms:modified xsi:type="dcterms:W3CDTF">2006-09-16T00:00:00Z</dcterms:modified></cp:coreProperties>`
)

func WriteExcel(filePath string, content [][]string) error {
	if err := Mkdir(filePath); err != nil {
		return err
	}
	file := excelize.NewFile()
	file.Pkg.Store(defaultXMLPathDocPropsCore, []byte(xml.Header+templateDocpropsCore))
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

func ReadExcel(filePath string, sheetName string) (content [][]string, err error) {
	file, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, err
	}
	defer func() {
		// Close the spreadsheet.
		if err := file.Close(); err != nil {
			log.Println(err)
		}
	}()
	return file.GetRows(sheetName)
}

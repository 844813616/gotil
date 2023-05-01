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
			if err := file.SetCellDefault(defaultSheetName, cellName, value); err != nil {
				return err
			}
		}
	}
	if err := file.SaveAs(filePath); err != nil {
		return err
	}
	return nil
}

func ReadExcel(filePath string, sheetName string) (content [][]string, err error) {
	file, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Println(err)
		}
	}()
	return file.GetRows(sheetName)
}

func UpdateExcel(file *excelize.File, sheetName string, cellNames []string, calAreas []Area, processor func(in [][][]string) string) error {
	index, err := file.GetSheetIndex(sheetName)
	if err != nil {
		return err
	}
	file.SetActiveSheet(index)
	for _, cellName := range cellNames {
		var in [][][]string
		for _, calAreas := range calAreas {
			cellNameArray := calAreas.ToCellNameArray(cellName)
			in = append(in, cellNameArray)
		}
		in = getAreasCellValue(file, sheetName, in)
		out := processor(in)
		if err := file.SetCellDefault(sheetName, cellName, out); err != nil {
			log.Println(err)
		}
	}
	if err := file.Save(); err != nil {
		return err
	}
	return nil
}

type Value interface {
	Get(int) int
}

type Absolute int

func (a Absolute) Get(_ int) int {
	return int(a)
}

type Relative int

func (r Relative) Get(relative int) int {
	return int(r) + relative
}

type Coordinate struct {
	row Value
	col Value
}

func (c Coordinate) GetCellName(relative string) string {
	col, row, err := excelize.CellNameToCoordinates(relative)
	if err != nil {
		return ""
	}
	cellName, err := excelize.CoordinatesToCellName(c.col.Get(col), c.row.Get(row))
	if err != nil {
		return ""
	}
	return cellName
}

type Area struct {
	leftUp      Coordinate
	rightBottom Coordinate
}

func (a Area) ToCellNameList(relative string) []string {
	left, right, up, bottom, err := a.GetBoundary(relative)
	if err != nil {
		return nil
	}
	var ret []string
	for col := left; col <= right; col++ {
		for row := up; row <= bottom; row++ {
			cellName, err := excelize.CoordinatesToCellName(col, row)
			if err != nil {
				return nil
			}
			ret = append(ret, cellName)
		}
	}
	return ret
}

func (a Area) ToCellNameArray(relative string) [][]string {
	left, right, up, bottom, err := a.GetBoundary(relative)
	if err != nil {
		return nil
	}
	var ret [][]string
	for row := up; row <= bottom; row++ {
		var rowRet []string
		for col := left; col <= right; col++ {
			cellName, err := excelize.CoordinatesToCellName(col, row)
			if err != nil {
				return nil
			}
			rowRet = append(rowRet, cellName)
		}
		ret = append(ret, rowRet)
	}
	return ret
}

func (a Area) GetBoundary(relative string) (left, right, up, bottom int, err error) {
	leftUpCellName := a.leftUp.GetCellName(relative)
	left, up, err = excelize.CellNameToCoordinates(leftUpCellName)
	if err != nil {
		return 0, 0, 0, 0, err
	}
	rightBottomCellName := a.rightBottom.GetCellName(relative)
	right, bottom, err = excelize.CellNameToCoordinates(rightBottomCellName)
	if err != nil {
		return 0, 0, 0, 0, err
	}
	return left, right, up, bottom, nil
}

func getAreasCellValue(file *excelize.File, sheetName string, cellNames [][][]string) [][][]string {
	var ret [][][]string
	for _, area := range cellNames {
		var retArea [][]string
		for _, row := range area {
			var retRow []string
			for _, col := range row {
				cellName := col
				value, err := file.GetCellValue(sheetName, cellName)
				if err != nil {
					return nil
				}
				retRow = append(retRow, value)
			}
			retArea = append(retArea, retRow)
		}
		ret = append(ret, retArea)
	}
	return ret
}

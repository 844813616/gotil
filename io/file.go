package io

import (
	"bufio"
	"log"
	"os"
	"path"
	"strings"
)

func ReadFile(filePath string) (ret []string, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println(err)
		}
	}(file)
	ret = []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		text = strings.TrimSpace(text)
		ret = append(ret, text)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return ret, err
}

func WriteFile(filePath string, content []string) (err error) {
	dir := path.Dir(filePath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println(err)
		}
	}(file)
	writer := bufio.NewWriter(file)
	for _, v := range content {
		_, err := writer.WriteString(v + "\n")
		if err != nil {
			return err
		}
	}
	if err = writer.Flush(); err != nil {
		return err
	}
	return nil
}

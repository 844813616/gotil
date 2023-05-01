package io

import (
	"encoding/json"
	"io"
)

func ReadJson(filePath string) (map[string]interface{}, error) {
	file, err := OpenFile(filePath, READ)
	if err != nil {
		return nil, err
	}
	all, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	var m map[string]interface{}
	if err := json.Unmarshal(all, &m); err != nil {
		return nil, err
	}
	return m, nil
}

func WriteJson(filePath string, m map[string]interface{}) error {
	file, err := OpenFile(filePath, WRITE)
	if err != nil {
		return err
	}
	marshal, err := json.Marshal(m)
	if err != nil {
		return err
	}
	_, err = io.WriteString(file, string(marshal))
	if err != nil {
		return err
	}
	return nil
}

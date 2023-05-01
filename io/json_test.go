package io

import (
	"encoding/json"
	"log"
	"testing"
)

var testJsonPath = "../temp/json_test.json"

func TestReadJson(t *testing.T) {
	j, err := ReadJson(testJsonPath)
	if err != nil {
		log.Println(err)
		return
	}
	marshal, err := json.Marshal(j)
	if err != nil {
		return
	}
	println(string(marshal))
}

func TestWriteJson(t *testing.T) {
	err := WriteJson(testJsonPath, map[string]interface{}{"1": 1, "2": "2", "3": []int{1, 2, 3}})
	if err != nil {
		log.Println(err)
		return
	}
}

func TestName(t *testing.T) {
	j := `[{"1":[1,2]}, {"2":"2"}]`
	var l []map[string]interface{}
	err := json.Unmarshal([]byte(j), &l)
	if err != nil {
		return
	}
	marshal, err := json.Marshal(l)
	if err != nil {
		return
	}
	println(string(marshal))
}

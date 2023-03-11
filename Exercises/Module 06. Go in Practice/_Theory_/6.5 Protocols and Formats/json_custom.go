package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type strslice []string

func (ss *strslice) UnmarshalJSON(data []byte) error { //для специфических полей JSON можно определить кастомный анмаршалинг
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	} //к нам придет массив байт, который надо анмаршалить
	*ss = strings.Split(s, ",") // готово, разбили строку по запятой
	return nil
}


type RequestContentTag struct {
	User    string
	Message string `json:"msg"` //json пакет как и остальные маршалеры поддерживает кастомные имена полей
	Tags strslice `json:"tags"`
}

type RequestTagged struct {
	Request RequestContentTag
	Author  string `json:"user"`
}

func LoadAndParseJsonToCustomisedStruct() {
	jsonData, err := os.ReadFile("examples/example_custom.json")
	fmt.Println(err)
	var request RequestTagged
	fmt.Println(json.Unmarshal(jsonData, &request))
	fmt.Println(request)
}
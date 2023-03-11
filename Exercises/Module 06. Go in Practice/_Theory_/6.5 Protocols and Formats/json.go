package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type RequestContent struct {
	User    string
	Message string `json:"msg"`
}

type Request struct {
	Request RequestContent
	Author  string `json:"user"`
}

func LoadAndParseJson() {
	jsonData, err := os.ReadFile("examples/example.json")
	fmt.Println(err)
	var request Request
	fmt.Println(json.Unmarshal(jsonData, &request))
	fmt.Println(request)
}

func LoadAndParseRawMsgToMap() {
	jsonData, _ := os.ReadFile("examples/example.json")
	var objmap map[string]interface{} //если мы не знаем точно что в json - самый простой способ разбить его в map[string]interface{}
	json.Unmarshal(jsonData, &objmap)
	fmt.Println(objmap)
}


func LoadAndParseRawMsg() {
	jsonData, _ := os.ReadFile("examples/example.json")
	var objmap map[string]json.RawMessage //для анмаршалинга только специфических полей мы можем воспользоваться json.RawMessage
	json.Unmarshal(jsonData, &objmap)
	fmt.Println(objmap)
	var internalMap map[string]string
	json.Unmarshal(objmap["request"], &internalMap)
	fmt.Println(internalMap)
}


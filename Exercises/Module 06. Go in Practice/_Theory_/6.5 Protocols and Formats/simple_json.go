package main

import (
	"encoding/json"
	"fmt"
)

//Определим структуры для хранения json
type Dimensions struct {
	Height int
	Width  int
}

type Bird struct {
	Species     string
	Description string
	Dimensions  Dimensions
}

func ParseJson() {
	birdJson := `{"species":"pigeon","description":"likes to perch on rocks", "dimensions":{"height":24,"width":10}}`
	var bird Bird
	err := json.Unmarshal([]byte(birdJson), &bird) //анмаршалим нашу строку в заранее обьявленную структуру

	if err != nil {
		panic(err)
	}

	fmt.Println(bird)
}

func CreateJson() {
	bird := Bird{
		Species:     "Eagle",
		Description: "Cool eagle",
		Dimensions: Dimensions{
			Height: 100,
			Width:  50,
		},
	}
	data, _ := json.MarshalIndent(bird,"", "  ") //Маршалим с отступами, чтобы было удобно читать, префикс не нужен, поэтому второй параметр - пустая строка
	fmt.Println(string(data))
}

func ParseJsonArrays() {
	arraysJson := `["one","two","three"]`

	var nums []string

	json.Unmarshal([]byte(arraysJson), &nums) //JSON разбирается в слайсы так же как и в обьекты, нужно только заранее знать что там будет

	fmt.Println(nums)
	fmt.Println(nums[0])
}
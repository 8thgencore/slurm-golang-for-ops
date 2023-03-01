package main

import "fmt"

func main() {

	sl := []int64{9, 8, 7}
	for key, value := range sl {
		fmt.Printf("key: %v, val: %v \n", key, value)
	}
	//key: 0, val: 9
	//key: 1, val: 8
	//key: 2, val: 7

	for _, value := range sl {
		fmt.Println(value)
	}
	//9
	//8
	//7

	ages := map[string]int{
		"Андрей":    30,
		"Анастасия": 25,
	}
	for key, value := range ages {
		fmt.Println(key)
		fmt.Println(value)
	}
	//Андрей
	//30
	//Анастасия
	//25

	word := "слёрм"

	for i := 0; i < len(word); i++ {
		fmt.Println(word[i])
		fmt.Printf("%T", word[i])
	}
	//209
	//uint8
	//
	//129
	//uint8
	//
	//208
	//uint8
	//
	//187
	//uint8
	//
	//209
	//uint8
	//
	//145
	//uint8
	//
	//209
	//uint8
	//
	//128
	//uint8
	//
	//208
	//uint8
	//
	//188
	//uint8

	for key, value := range word {
		fmt.Println(key)
		fmt.Println(value)
		fmt.Printf("%T", value)
	}
	//0
	//1089
	//int32
	//
	//2
	//1083
	//int32
	//
	//4
	//1105
	//int32
	//
	//6
	//1088
	//int32
	//
	//8
	//1084
	//int32
}

package main

import "fmt"

func main() {
	var m1 map[int32]bool
	var m2 map[string]string

	m3 := make(map[int]int)

	ages := map[string]int{
		"Андрей":    30,
		"Анастасия": 25,
		// ...
	}
	age := ages["Андрей"]       // 30
	ages["Наталья"] = 31        // [Анастасия:25 Андрей:30 Наталья:31]
	fmt.Println(ages["Михаил"]) // 0
	//ages["Михаил"]++ // [Анастасия:25 Андрей:30 Наталья:31 Михаил:1]
	ages["Михаил"] = ages["Михаил"] + 1 // [Анастасия:25 Андрей:30 Наталья:31 Михаил:1]

	fmt.Println(m1, m2, m3, age)
}

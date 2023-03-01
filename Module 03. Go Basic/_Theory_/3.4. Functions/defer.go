package main

import "fmt"

func end() {
	fmt.Println("function ends")
}

func main() {
	defer end()
	num := 5

	defer func(x int) {
		fmt.Println(x)
	}(num)
	num = 20

	//5
	//function ends
}

package main

import "fmt"

func main() {
	type Point struct {
		X int
		Y int
		//Z string
		//T []int64
	}

	p := Point{
		X: 5,
		Y: 19,
	}

	p = Point{5, 15}

	fmt.Println(p.X) // 5
	fmt.Println(p.Y) // 15

	p.X = 55
	fmt.Println(p.X) // 55

}

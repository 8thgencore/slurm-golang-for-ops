package main

import "fmt"

func foo() {
	defer func() {
		err := recover()
		fmt.Printf("Error: %v\n", err)
		fmt.Printf("Recovery\n")
	}()
	panic("This is panic")
}

func main() {
	foo()
}

package main

import "fmt"

const panicMessage = "panic happened"

func foo() {
	defer func() {
		err := recover()
		fmt.Printf("Error: %v\n", err)
		fmt.Printf("Recovery\n")
	}()
	panic(panicMessage)
}

func main() {
	foo()
}

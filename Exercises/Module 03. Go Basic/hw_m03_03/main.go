package main

import (
	"fmt"
)

func isSorted(words []string) bool {
	for i := 1; i < len(words); i++ {
		if words[i] < words[i-1] {
			return false
		}
	}
	return true
}

func main() {
	words := []string{"Close", "Danke", "Ferder", "Liben"}
	fmt.Println(isSorted(words))
}

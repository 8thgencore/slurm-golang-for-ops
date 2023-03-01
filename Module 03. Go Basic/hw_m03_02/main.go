package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateArray() []int {
	rand.Seed(time.Now().UnixNano())

	array := make([]int, 10)
	for i := 0; i < 10; i++ {
		var n int
		n = rand.Intn(31)
		array[i] = n
	}
	return array
}

func hasDuplicate(nums []int) bool {
	used := make(map[int]bool)
	for _, num := range nums {
		if used[num] {
			return true
		}
		used[num] = true

	}
	return false
}

func main() {
	nums := generateArray()
	fmt.Printf("Array %v\n", nums)
	fmt.Println(hasDuplicate(nums))
}

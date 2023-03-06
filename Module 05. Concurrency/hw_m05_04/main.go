package main

import (
	"fmt"
	"sync"
)

func firstArray(ch chan<- int, wg *sync.WaitGroup) {
	a := []int{1, 3, 5, 7, 9}
	for _, v := range a {
		ch <- v
	}
	wg.Done()
}

func secondArray(ch chan<- int, wg *sync.WaitGroup) {
	b := []int{2, 4, 6, 8, 10}
	for _, v := range b {
		ch <- v
	}
	wg.Done()
}

func main() {
	wg := &sync.WaitGroup{}

	var c [10]int

	ch1 := make(chan int)
	ch2 := make(chan int)

	go firstArray(ch1, wg)
	wg.Add(1)

	go secondArray(ch2, wg)
	wg.Add(1)

	for i:=0; i<10; i++ {
		c[i] = <- ch1
		i++
		c[i] = <- ch2
	}

	fmt.Printf("Result: %v\n", c)
}

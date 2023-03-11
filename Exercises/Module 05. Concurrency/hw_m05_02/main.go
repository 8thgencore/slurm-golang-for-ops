package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

const timeout = time.Second * 2

func monitoring(t *time.Timer, quit chan int, ans chan error) {
	for {
		select {
		case <-t.C:
			ans <- errors.New("Time is expected")
			return
		case <-quit:
			ans <- nil
			return
		}
	}
}

func work(wg *sync.WaitGroup, num int) {
	time.Sleep(2 * time.Second)
	fmt.Printf("This is goroutine: %v\n", num)
	wg.Done()
}

func foo() error {
	wg := &sync.WaitGroup{}
	timer := time.NewTimer(timeout)
	quit := make(chan int, 1)
	ans := make(chan error)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go work(wg, i)
	}

	go monitoring(timer, quit, ans)

	wg.Wait()
	quit <- 1

	return <-ans
}

func main() {
	err := foo()
	if err != nil {
		fmt.Printf("Error: %v", err)
	} else {
		fmt.Println("All goroutines were completed")
	}
}

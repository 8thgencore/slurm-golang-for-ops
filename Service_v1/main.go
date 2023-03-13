package main

import (
	"context"
	"fmt"
	"hw_m04_03/gometr"
	"time"
)

const workTime = 30

func catchPanic() {
    if r := recover(); r != nil { // проверяем, есть ли паника
        fmt.Println("Recovered from", r) // выводим значение паники
    }
}


func main() {
	defer catchPanic()

	che := new(gometr.Checker)
	ch := make(chan gometr.Checkable, 1)

	parent := context.Background()
	_, cancel := context.WithCancel(parent)

	// add object with channel
	ch <- gometr.NewGoMetrClient("1/dev/req", 2)
	go che.Add(ch)
	ch <- gometr.NewGoMetrClient("2/dev/req", 8)
	go che.Add(ch)
	ch <- gometr.NewGoMetrClient("3/dev/req", 14)
	go che.Add(ch)
	ch <- gometr.NewGoMetrClient("4/dev/req", 20)
	go che.Add(ch)
	ch <- gometr.NewGoMetrClient("5/dev/req", 26)
	go che.Add(ch)

	// run checker
	go che.Run()
	time.Sleep((workTime / 6) * time.Second)

	// add object with channel
	ch <- gometr.NewGoMetrClient("7/dev/req", 12)
	go che.Add(ch)
	time.Sleep(workTime * time.Second)

	// panic
	panic("This is panic")

	// stop checker
	che.Stop(cancel)
}

package gometr

import (
	"context"
	"fmt"
	"time"
)

type Measurable interface {
	GetMetrics() string
}

type Checkable interface {
	Measurable
	Ping() error
	GetID() string
	Health() bool
}

// --------------------------------------------------
// Checker struct
// --------------------------------------------------
type Checker struct {
	items []Checkable
}

// add object with channel
func (c *Checker) Add(ch <-chan Checkable) {
	c.items = append(c.items, <-ch)
}

func (c *Checker) String() string {
	var str string
	for _, val := range c.items {
		str += val.GetID()
		str += "\n"
	}
	return str
}

func (c *Checker) Check() {
	for _, val := range c.items {
		if !val.Health() {
			fmt.Println(val.GetID() + " не работает")
		}
	}
}

func (c *Checker) Run() {
	fmt.Printf("‼️ Проверки запущены\n")
	ticker := time.NewTicker(5 * time.Second)

	for tick := range ticker.C {
		fmt.Println("")
		for _, val := range c.items {
			go check(tick, val)
		}
	}
}

func (c *Checker) Stop(cancel context.CancelFunc) {
	fmt.Printf("‼️ Проверки остановлены\n")
	cancel()
}

func check(tick time.Time, client Checkable) {
	if !client.Health() {
		fmt.Printf("%s не работает, время: %v\n", client.GetID(), tick)
	}
}

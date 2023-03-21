package checker

import (
	"context"
	log "service/pkg/logger"
)

type Measurable interface {
	GetMetrics() error
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
			log.Errorf(val.GetID() + " не работает")
		}
	}
}

func (c *Checker) Run() {
	log.Infof("‼️ Запуск проверок")
	for _, val := range c.items {
		go check(val)
	}
}

func (c *Checker) Stop(cancel context.CancelFunc) {
	log.Infof("‼️ Проверки остановлены")
	cancel()
}

func check(client Checkable) {
	if !client.Health() {
		log.Warnf("%s не работает, время: %v\n", client.GetID())
	} else {
		err := client.GetMetrics()
		if err != nil {
			log.Warnf(err.Error())
		}
	}
}

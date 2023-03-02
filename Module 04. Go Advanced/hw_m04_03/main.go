package main

import (
	"hw_m04_03/gometr"
	"time"
)

func main() {
	che := new(gometr.Checker)

	che.Add(gometr.NewGoMetrClient("1/dev/req", 120))
	che.Add(gometr.NewGoMetrClient("2/dev/sek", 140))
	che.Add(gometr.NewGoMetrClient("3/dev/cat", 160))
	che.Add(gometr.NewGoMetrClient("4/dev/vim", 180))
	che.Add(gometr.NewGoMetrClient("5/dev/duf", 200))

	go che.Run()
	time.Sleep(5 * time.Second)

	che.Add(gometr.NewGoMetrClient("6/dev/fuf", 220))
	time.Sleep(30 * time.Second)
}

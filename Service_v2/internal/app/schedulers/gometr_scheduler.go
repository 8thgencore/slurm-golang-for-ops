package schedulers

import (
	"bufio"
	"fmt"
	"net/http"
	"service/internal/app/processors"
	cfg "service/internal/config"

	log "github.com/sirupsen/logrus"
)

type GometrScheduler struct {
	processor *processors.MetricsProcessor
}

func NewGometrScheduler(processor *processors.MetricsProcessor) *GometrScheduler {
	scheduler := new(GometrScheduler)
	scheduler.processor = processor
	return scheduler
}

func (scheduler *GometrScheduler) ParseGometr() error {
	log.Println("[GometrScheduler] Start ParseGometr")
	log.Println(cfg.ExternalConfig.GometrUrl + "/metrics")

	resp, err := http.Get(cfg.ExternalConfig.GometrUrl  + "/metrics")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

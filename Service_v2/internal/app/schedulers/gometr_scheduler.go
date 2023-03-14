package schedulers

import (
	"bufio"
	"net/http"
	"service/internal/app/models"
	"service/internal/app/processors"
	cfg "service/internal/config"
	log "service/pkg/logger"
	"strings"
	"time"
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
	log.Infof("[GometrScheduler] Start ParseGometr")

	resp, err := http.Get(cfg.ExternalConfig.GometrUrl + "/metrics")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	currentDate := time.Now()
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		text := scanner.Text()
		if string(text[0]) == "#" {
			continue
		}
		splitStr := strings.Split(text, " ")
		scheduler.processor.Add(models.Metric{
			Name:  splitStr[0],
			Value: splitStr[1],
			Date:  currentDate,
		})
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

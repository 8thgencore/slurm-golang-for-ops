package gometrclient

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"service/internal/app/models"
	"service/internal/app/processors"
	cfg "service/internal/config"
	log "service/pkg/logger"
	"strings"
	"time"
)

type GoMetrClient struct {
	processor *processors.MetricsProcessor
}

func NewGoMetrClient(processor *processors.MetricsProcessor) *GoMetrClient {
	checker := new(GoMetrClient)
	checker.processor = processor

	return checker
}

func (g *GoMetrClient) GetMetrics() error {
	resp, err := http.Get(cfg.ExternalConfig.GometrURL + "/metrics")
	if err != nil {
		log.Errorf("Cannot get metrics")
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
		g.processor.Add(models.Metric{
			Name:  splitStr[0],
			Value: splitStr[1],
			Date:  currentDate,
		})
	}

	return scanner.Err()
}

func (g GoMetrClient) Ping() error {
	return nil
}

func (g GoMetrClient) GetID() string {
	return ""
}

func (g GoMetrClient) Health() bool {
	resp, err := http.Get(cfg.ExternalConfig.GometrURL + "/health")
	if err != nil {
		log.Errorf("Cannot get metrics")
		return false
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("Cannot get metrics")
	}

	var health Health
	err = json.Unmarshal(body, &health)
	if err != nil {
		log.Errorf("Cannot get metrics")
	}

	if health.Checks.PingMysql.Status == string(PassStatus) {
		return true
	}
	return false
}

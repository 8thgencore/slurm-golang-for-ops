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

type Client struct {
	processor *processors.MetricsProcessor
}

func NewClient(processor *processors.MetricsProcessor) *Client {
	checker := new(Client)
	checker.processor = processor

	return checker
}

func (c *Client) GetMetrics() error {
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
		c.processor.Add(models.Metric{
			Name:  splitStr[0],
			Value: splitStr[1],
			Date:  currentDate,
		})
	}

	return scanner.Err()
}

func (c *Client) Ping() error {
	return nil
}

func (c *Client) GetID() string {
	return "gometr_service"
}

func (c *Client) Health() bool {
	resp, err := http.Get(cfg.ExternalConfig.GometrURL + "/health")
	if err != nil {
		log.Errorf("Service is not healthy")
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

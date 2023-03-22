package googlemetr

import (
	"service/internal/app/models"
	"service/internal/app/processors"
	log "service/pkg/logger"
	"time"

	"github.com/go-ping/ping"
)

const host = "www.google.com"

type Client struct {
	processor *processors.MetricsProcessor
}

func NewClient(processor *processors.MetricsProcessor) *Client {
	checker := new(Client)
	checker.processor = processor

	return checker
}

func (c *Client) GetMetrics() error {
	return nil
}

func (c *Client) Ping() error {
	var err error
	pinger, err := ping.NewPinger(host)
	if err != nil {
		log.Errorf("Dont create Pinger for %s", c.GetID())
		return err
	}
	pinger.Count = 3
	err = pinger.Run()
	if err != nil {
		log.Errorf("%s not responding. %v", c.GetID(), err)
		return err
	}

	return nil
}

func (c *Client) GetID() string {
	return host
}

func (c *Client) Health() bool {
	err := c.Ping()
	currentDate := time.Now()

	if err != nil {
		c.processor.Add(models.Metric{
			Name:  c.GetID(),
			Value: "unavailable",
			Date:  currentDate,
		})
		log.Errorf("Server not pinger")

		return false
	}

	c.processor.Add(models.Metric{
		Name:  c.GetID(),
		Value: "available",
		Date:  currentDate,
	})

	return true
}

package gometr

import (
	"time"
)

// --------------------------------------------------
// HealtCheck struct
// --------------------------------------------------
type ErrorType string

const (
	PassStatus ErrorType = "pass"
	FailStatus ErrorType = "fail"
)

type HealthCheck struct {
	ServiceID string
	Status    ErrorType
}

// --------------------------------------------------
// GoMetrClient struct
// --------------------------------------------------
var startTime = time.Now()

type GoMetrClient struct {
	url     string
	timeout int
}

func NewGoMetrClient(url string, timeout int) GoMetrClient {
	return GoMetrClient{
		url,
		timeout,
	}
}

func (g GoMetrClient) GetMetrics() string {
	return string(rune(g.timeout))
}

func (g GoMetrClient) Ping() error {
	return nil
}

func (g GoMetrClient) GetID() string {
	return g.url
}

func (g GoMetrClient) Health() bool {
	return g.getHealth().Status == PassStatus
}

func (g GoMetrClient) getHealth() HealthCheck {
	timeout := startTime.Add(time.Duration(g.timeout) * time.Second)
	end := time.Now().After(timeout)

	if end {
		return HealthCheck{g.url, FailStatus}
	}
	return HealthCheck{g.url, PassStatus}
}

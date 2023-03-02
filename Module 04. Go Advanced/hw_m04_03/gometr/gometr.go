package gometr

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
	return g.GetHealth().Status == PassStatus
}

func (g GoMetrClient) GetHealth() HealthCheck {
	if g.url[0]%2 == 0 {
		return HealthCheck{g.url, PassStatus}
	} else {
		return HealthCheck{g.url, FailStatus}
	}
}

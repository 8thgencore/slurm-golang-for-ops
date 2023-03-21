package gometrclient

// --------------------------------------------------
// HealthCheck struct
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
// Health struct
// --------------------------------------------------
type Health struct {
	Status    string      `json:"status"`
	ServiceID string      `json:"service_id"`
	Checks    ServiceList `json:"checks"`
}

type ServiceList struct {
	PingMysql Service `json:"ping_mysql"`
}

type Service struct {
	ComponentID   string `json:"component_id"`
	ComponentType string `json:"component_type"`
	Status        string `json:"status"`
}

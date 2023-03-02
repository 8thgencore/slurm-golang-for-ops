package models

type CheckStatus string

const (
	CheckStatusPass CheckStatus = "pass"
	CheckStatusWarn CheckStatus = "warn"
	CheckStatusFail CheckStatus = "fail"
)

type CheckResponse struct {
	Status     CheckStatus `json:"status"`
	ServiceID  string      `json:"service_id,omitempty"`
	Checks     Checks      `json:"checks,omitempty"`
}

type Checks map[string]CheckResult

type CheckResult struct {
	ComponentID       string      `json:"component_id"`
	ComponentType     string      `json:"component_type"`
	Status            CheckStatus `json:"status"`
}

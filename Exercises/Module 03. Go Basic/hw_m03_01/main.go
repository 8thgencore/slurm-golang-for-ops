package main

import (
	"fmt"
)

const (
	PassStatus = "pass"
	FailStatus = "fail"
)

type HealthCheck struct {
	ServiceID string
	Status    string
}

func GenerateCheck() []HealthCheck {
	var checks []HealthCheck
	for i := 0; i < 5; i++ {
		status := FailStatus
		if i%2 == 0 {
			status = PassStatus
		}
		check := HealthCheck{
			ServiceID: fmt.Sprintf("Service-%d", i),
			Status:    status,
		}
		checks = append(checks, check)

	}
	return checks
}

func main() {
	checks := GenerateCheck()
	for _, check := range checks {
		if check.Status == PassStatus {
			fmt.Println(check.ServiceID)
		}
	}
}

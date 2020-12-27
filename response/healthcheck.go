package response

import "time"

type HealthCheckResponse struct {
	HealthStatus string    `json:"health_status"`
	DBTimestamp  time.Time `json:"database_timestamp"`
	Environment  string    `json:"environment"`
}

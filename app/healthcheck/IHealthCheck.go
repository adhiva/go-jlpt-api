package healthcheck

import "go-jlpt-n5/models"

type IHealthCheckUsecase interface {
	GetDBTimestamp() models.HealthCheck
}

type IHealthCheckRepository interface {
	GetDBTimestamp() models.HealthCheck
}

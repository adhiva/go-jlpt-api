package usecase

import (
	HealthCheckInterface "go-jlpt-n5/app/healthcheck"
	"go-jlpt-n5/models"
)

type HealthCheckUsecase struct {
	HealthCheckRepository HealthCheckInterface.IHealthCheckRepository
}

func NewHealthCheckUsecase(hcr HealthCheckInterface.IHealthCheckRepository) HealthCheckInterface.IHealthCheckUsecase {
	return &HealthCheckUsecase{
		HealthCheckRepository: hcr,
	}
}

func (a *HealthCheckUsecase) GetDBTimestamp() models.HealthCheck {
	healthCheck := a.HealthCheckRepository.GetDBTimestamp()
	return healthCheck
}

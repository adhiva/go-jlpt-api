package repository

import (
	HealthCheckInterface "go-jlpt-n5/app/healthcheck"
	"go-jlpt-n5/models"

	"gorm.io/gorm"
)

type HealthCheckRepository struct {
	Conn *gorm.DB
}

func NewHealthCheckRepository(Conn *gorm.DB) HealthCheckInterface.IHealthCheckRepository {
	return &HealthCheckRepository{Conn}
}

func (m *HealthCheckRepository) GetDBTimestamp() models.HealthCheck {
	var healthCheck models.HealthCheck

	tx := m.Conn.Begin()
	tx.Raw("SELECT current_timestamp").Scan(&healthCheck)
	tx.Commit()

	return healthCheck
}

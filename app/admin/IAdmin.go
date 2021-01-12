package admin

import "go-jlpt-n5/models"

type IAdminUsecase interface {
	Get() ([]*models.Admin, error)
	Store() (*models.Admin, error)
}

type IAdminRepository interface {
	GetDBTimestamp() models.HealthCheck
	Get() ([]*models.Admin, error)
}

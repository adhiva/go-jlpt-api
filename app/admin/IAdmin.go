package admin

import (
	"go-jlpt-n5/models"
	"go-jlpt-n5/request"
)

type IAdminUsecase interface {
	Get() ([]*models.Admin, error)
	Store(*request.Admin) (*models.Admin, error)
}

type IAdminRepository interface {
	GetDBTimestamp() models.HealthCheck
	Get() ([]*models.Admin, error)
}

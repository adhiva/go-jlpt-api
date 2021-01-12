package usecase

import (
	AdminInterface "go-jlpt-n5/app/admin"
	"go-jlpt-n5/models"
)

type AdminUsecase struct {
	AdminRepository AdminInterface.IAdminRepository
}

func NewAdminUsecase(ar AdminInterface.IAdminRepository) AdminInterface.IAdminUsecase {
	return &AdminUsecase{
		AdminRepository: ar,
	}
}

func (a *AdminUsecase) GetDBTimestamp() models.HealthCheck {
	healthCheck := a.AdminRepository.GetDBTimestamp()
	return healthCheck
}

func (a *AdminUsecase) Get() ([]*models.Admin, error) {
	admins, err := a.AdminRepository.Get()
	if err != nil {
		return nil, err
	}
	return admins, nil
}

func (a *AdminUsecase) Store() (*models.Admin, error) {
	return &models.Admin{}, nil
}

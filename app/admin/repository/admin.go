package repository

import (
	"fmt"
	AdminInterface "go-jlpt-n5/app/admin"
	"go-jlpt-n5/models"

	"gorm.io/gorm"
)

type AdminRepository struct {
	Conn *gorm.DB
}

func NewAdminRepository(Conn *gorm.DB) AdminInterface.IAdminRepository {
	return &AdminRepository{Conn}
}

func (m *AdminRepository) GetDBTimestamp() models.HealthCheck {
	var healthCheck models.HealthCheck

	tx := m.Conn.Begin()
	tx.Raw("SELECT current_timestamp").Scan(&healthCheck)
	tx.Commit()

	return healthCheck
}

func (m *AdminRepository) Get() ([]*models.Admin, error) {
	var (
		admins []*models.Admin
		err    error
	)

	tx := m.Conn.Begin()
	txCount := tx

	if err = tx.Find(&admins).Error; err != nil {
		tx.Rollback()
		txCount.Rollback()
		return admins, err
	}
	tx.Commit()
	txCount.Commit()
	fmt.Println("Check Admin : ", admins)

	return admins, nil
}

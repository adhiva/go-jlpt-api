package models

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	ID                     int       `json:"id"`
	FirstName              string    `json:"first_name"`
	LastName               string    `json:"last_name"`
	Email                  string    `json:"email"`
	PhoneNumberCountryCode string    `json:"phone_number_country_code"`
	PhoneNumber            string    `json:"phone_number"`
	Country                string    `json:"country"`
	CreatedAt              time.Time `json:"created_at"`
	UpdatedAt              time.Time `json:"updated_at"`
}

type AdminToken struct {
	Token string `json:"token"`
}

func ActiveAdmin(db *gorm.DB) *gorm.DB {
	return db.Where("is_active = ?", true)
}

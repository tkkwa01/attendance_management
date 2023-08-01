package domain

import "gorm.io/gorm"

type Employees struct {
	gorm.Model
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

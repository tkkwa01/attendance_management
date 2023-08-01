package domain

import "gorm.io/gorm"

type SalaryTypes struct {
	gorm.Model
	ID   uint   `json:"id"`
	Type string `json:"type"`
}

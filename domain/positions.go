package domain

import "gorm.io/gorm"

type Positions struct {
	gorm.Model
	ID   uint   `json:"id"`
	Type string `json:"type"`
}

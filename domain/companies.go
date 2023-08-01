package domain

import "gorm.io/gorm"

type Companies struct {
	gorm.Model
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

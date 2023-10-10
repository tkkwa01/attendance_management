package domain

import (
	"attendance-management/packages/context"
	"attendance-management/packages/validation"
	"attendance-management/resource/request"
	"gorm.io/gorm"
)

type SalaryTypes struct {
	gorm.Model
	ID               uint   `json:"id"`
	Type             string `json:"type"`
	SalaryTypeNumber uint   `json:"salary_type_number" gorm:"unique"`
}

func NewSalaryType(ctx context.Context, req *request.CreateSalaryType) (*SalaryTypes, error) {
	salaryType := &SalaryTypes{
		Type: req.Type,
	}
	err := validation.Validate().Struct(salaryType)
	if err != nil {
		return nil, err
	}

	return salaryType, nil
}

package domain

import (
	"attendance-management/packages/context"
	"attendance-management/packages/validate"
	"attendance-management/resource/request"
	"gorm.io/gorm"
)

type SalaryTypes struct {
	gorm.Model
	ID               uint   `json:"id"`
	Type             string `json:"type"`
	SalaryTypeNumber uint   `json:"salary_type_number"`
}

func NewSalaryType(ctx context.Context, req *request.CreateSalaryType) (*SalaryTypes, error) {
	salaryType := &SalaryTypes{
		Type: req.Type,
	}
	err := validate.Validate(salaryType)
	if err != nil {
		return nil, err
	}

	return salaryType, nil
}

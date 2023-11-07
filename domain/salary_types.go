package domain

import (
	"attendance-management/packages/context"
	"attendance-management/packages/validation"
	"attendance-management/resource/request"
)

type SalaryTypes struct {
	ID               uint          `json:"id" gorm:"primaryKey;autoIncrement"`
	Type             string        `json:"type" gorm:"type:varchar(255);not null"`
	SalaryTypeNumber uint          `json:"salary_type_number" gorm:"unique"`
	Employments      []Employments `gorm:"foreignKey:SalaryTypeID"`
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

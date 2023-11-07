package domain

import (
	"attendance-management/packages/context"
	"attendance-management/resource/request"
)

type SalaryTypes struct {
	ID               uint          `json:"id" gorm:"primaryKey;autoIncrement"`
	Type             string        `json:"type" gorm:"type:varchar(255);not null"`
	SalaryTypeNumber uint          `json:"salary_type_number" gorm:"unique"`
	Employments      []Employments `gorm:"foreignKey:SalaryTypeID"`
}

func NewSalaryType(ctx context.Context, dto *request.CreateSalaryType) (*SalaryTypes, error) {
	salaryType := SalaryTypes{
		Type:             dto.Type,
		SalaryTypeNumber: dto.SalaryTypeNumber,
	}

	if ctx.IsInValid() {
		return nil, ctx.ValidationError()
	}

	return &salaryType, nil
}

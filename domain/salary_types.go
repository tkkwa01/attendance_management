package domain

import (
	"attendance-management/packages/context"
	"attendance-management/resource/request"
)

type SalaryTypes struct {
	ID          uint          `json:"id" gorm:"primaryKey;autoIncrement"`
	Type        string        `json:"type" gorm:"type:varchar(255);not null"`
	Employments []Employments `json:"employments" gorm:"foreignKey:SalaryTypeID"`
}

func NewSalaryType(ctx context.Context, dto *request.CreateSalaryType) (*SalaryTypes, error) {
	salaryType := SalaryTypes{
		Type: dto.Type,
	}

	if ctx.IsInValid() {
		return nil, ctx.ValidationError()
	}

	return &salaryType, nil
}

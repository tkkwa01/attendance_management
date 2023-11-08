package domain

import (
	"attendance-management/packages/context"
	"attendance-management/resource/request"
)

type Companies struct {
	ID            uint          `json:"id" gorm:"primaryKey;autoIncrement"`
	Name          string        `json:"name" gorm:"type:varchar(255);not null"`
	CompanyNumber uint          `json:"company_number" gorm:"unique"`
	Employments   []Employments `json:"employments"`
}

func NewCompany(ctx context.Context, dto *request.CreateCompany) (*Companies, error) {
	company := &Companies{
		Name:          dto.Name,
		CompanyNumber: dto.CompanyNumber,
	}

	if ctx.IsInValid() {
		return nil, ctx.ValidationError()
	}

	return company, nil
}

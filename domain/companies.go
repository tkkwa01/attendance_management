package domain

import (
	"attendance-management/packages/context"
	"attendance-management/packages/validation"
	"attendance-management/resource/request"
)

type Companies struct {
	ID            uint          `json:"id" gorm:"primaryKey;autoIncrement"`
	Name          string        `json:"name" gorm:"type:varchar(255);not null"`
	CompanyNumber uint          `json:"company_number" gorm:"unique"`
	Employments   []Employments `json:"employments" gorm:"foreignKey:CompanyID"`
}

func NewCompany(ctx context.Context, req *request.CreateCompany) (*Companies, error) {
	company := &Companies{
		Name:          req.Name,
		CompanyNumber: req.CompanyNumber,
	}
	err := validation.Validate().Struct(company)
	if err != nil {
		return nil, err
	}
	return company, nil
}

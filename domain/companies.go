package domain

import (
	"attendance-management/packages/context"
	"attendance-management/packages/validate"
	"attendance-management/resource/request"
	"gorm.io/gorm"
)

type Companies struct {
	gorm.Model
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	CompanyNumber uint   `json:"company_number" gorm:"unique"`
}

func NewCompany(ctx context.Context, req *request.CreateCompany) (*Companies, error) {
	company := &Companies{
		Name:          req.Name,
		CompanyNumber: req.CompanyNumber,
	}
	err := validate.Validate(company)
	if err != nil {
		return nil, err
	}
	return company, nil
}

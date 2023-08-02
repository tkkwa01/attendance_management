package domain

import (
	"attendance-management/packages/validate"
	"attendance-management/resource/request"
	"context"
	"gorm.io/gorm"
)

type Employees struct {
	gorm.Model
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

func NewEmployee(ctx context.Context, req *request.CreateEmployee) (*Employees, error) {
	employee := &Employees{
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
	}
	err := validate.Validate(employee)
	if err != nil {
		return nil, err
	}
	return employee, nil
}

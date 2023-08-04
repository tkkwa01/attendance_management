package domain

import (
	"attendance-management/packages/context"
	"attendance-management/packages/validate"
	"attendance-management/resource/request"
	"gorm.io/gorm"
)

type Employees struct {
	gorm.Model
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	PhoneNumber    string `json:"phone_number"`
	EmployeeNumber uint   `json:"employee_number"` //GetByIDとかのときのキー
}

func NewEmployee(ctx context.Context, req *request.CreateEmployee) (*Employees, error) {
	employee := &Employees{
		Name:           req.Name,
		PhoneNumber:    req.PhoneNumber,
		EmployeeNumber: req.EmployeeNumber,
	}
	err := validate.Validate(employee)
	if err != nil {
		return nil, err
	}
	return employee, nil
}

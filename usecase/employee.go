package usecase

import (
	"attendance-management/domain"
	"attendance-management/resource/request"
	"context"
)

type EmployeeInputPort interface {
	Create() error
	GetByID() error
	Update() error
	Delete() error
}

type EmployeeOutputPort interface{
	Create()error
	GetByID()error
	Update()error
	Delete()error
}

type EmployeeRepository interface{

}

type employee struct{
	outputPort	EmployeeOutputPort
	EmployeeRepo EmployeeRepository
}

func(e employee) Create(ctx context.Context, req *request.CreateEmployee) error {
	newEmployee, err := domain.NewEmployee(ctx, req)
	err =
}

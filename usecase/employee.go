package usecase

import (
	"attendance-management/domain"
	"attendance-management/packages/context"
	"attendance-management/resource/request"
)

type EmployeeInputPort interface {
	Create(ctx context.Context, req *request.CreateEmployee) error
	//GetByID() error
	//Update() error
	//Delete() error
}

type EmployeeOutputPort interface {
	Create(id uint) error
	GetByID() error
	Update() error
	Delete() error
}

type EmployeeRepository interface {
	Create(ctx context.Context, employee *domain.Employees) (uint, error)
	NumberExist(ctx context.Context, number uint) error
}

type employee struct {
	outputPort   EmployeeOutputPort
	EmployeeRepo EmployeeRepository
}

type EmployeeInputFactory func(outputPort EmployeeOutputPort) EmployeeInputPort

func NewEmployeeInputFactory(er EmployeeRepository) EmployeeInputFactory {
	return func(o EmployeeOutputPort) EmployeeInputPort {
		return &employee{
			outputPort:   o,
			EmployeeRepo: er,
		}
	}
}

func (e employee) Create(ctx context.Context, req *request.CreateEmployee) error {
	newEmployee, err := domain.NewEmployee(ctx, req)
	// req.EmployeeNumberをkeyにしてdogが存在するか確認
	err = e.EmployeeRepo.NumberExist(ctx, req.EmployeeNumber)
	if err == nil {
		return ctx.Error().BadRequest("employee already exist")
	}

	id, err := e.EmployeeRepo.Create(ctx, newEmployee)
	if err != nil {
		return err
	}

	return e.outputPort.Create(id)
}

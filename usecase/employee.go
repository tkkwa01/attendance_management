package usecase

import (
	"attendance-management/domain"
	"attendance-management/packages/context"
	"attendance-management/resource/request"
)

type EmployeeInputPort interface {
	Create(ctx context.Context, req *request.CreateEmployee) error
	GetByID(ctx context.Context, number uint) error
	Update(ctx context.Context, req *request.UpdateEmployee) error
	Delete(ctx context.Context, number uint) error
}

type EmployeeOutputPort interface {
	Create(id uint) error
	GetByID(res *domain.Employees) error
	Update(res *domain.Employees) error
	Delete() error
}

type EmployeeRepository interface {
	Create(ctx context.Context, employee *domain.Employees) (uint, error)
	GetByID(ctx context.Context, number uint) (*domain.Employees, error)
	Update(ctx context.Context, employee *domain.Employees) error
	NumberExist(ctx context.Context, number uint) error
	Delete(ctx context.Context, number uint) error
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
	// req.EmployeeNumberをkeyにしてemployeeが存在するか確認
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

// q:なぜ引数でcontextを持たせている？
// a:contextを持たせることで、contextの中にerrorを持たせることができる
// これにより、errorを返すことができる
// また、contextを持たせることで、contextの中にDBを持たせることができる
// これにより、DBを呼び出すことができる
// また、contextを持たせることで、contextの中にuidを持たせることができる
// これにより、uidを呼び出すことができる
// また、contextを持たせることで、contextの中にrequestIDを持たせることができる
// これにより、requestIDを呼び出すことができる

// make getByID code
func (e employee) GetByID(ctx context.Context, number uint) error {
	// repositoryのGetByIDを呼び出す
	res, err := e.EmployeeRepo.GetByID(ctx, number)
	if err != nil {
		return err
	}
	return e.outputPort.GetByID(res)
}

func (e employee) Update(ctx context.Context, req *request.UpdateEmployee) error {
	employee, err := e.EmployeeRepo.GetByID(ctx, req.EmployeeNumber)
	if err != nil {
		return err
	}

	if req.Name != "" {
		employee.Name = req.Name
	}
	if req.PhoneNumber != "" {
		employee.PhoneNumber = req.PhoneNumber
	}

	err = e.EmployeeRepo.Update(ctx, employee)
	if err != nil {
		return err
	}
	return e.outputPort.Update(employee)
}

func (e employee) Delete(ctx context.Context, number uint) error {
	err := e.EmployeeRepo.Delete(ctx, number)
	if err != nil {
		return err
	}
	return e.outputPort.Delete()
}

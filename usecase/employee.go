package usecase

import (
	"attendance-management/config"
	"attendance-management/domain"
	"attendance-management/packages/context"
	"attendance-management/packages/errors"
	"attendance-management/resource/mail_body"
	"attendance-management/resource/request"
	"attendance-management/resource/response"
	"net/http"
)

type EmployeeInputPort interface {
	Create(ctx context.Context, req *request.EmployeeCreate) error
	GetByID(ctx context.Context, number uint) error
	Update(ctx context.Context, req *request.EmployeeUpdate) error
	Delete(ctx context.Context, number uint) error
	ResetPasswordRequest(ctx context.Context, req *request.EmployeeResetPasswordRequest) error
	ResetPassword(ctx context.Context, req *request.EmployeeResetPassword) error
	Login(ctx context.Context, req *request.EmployeeLogin) error
	RefreshToken(req *request.EmployeeRefreshToken) error
}

type EmployeeOutputPort interface {
	Create(id uint) error
	GetByID(res *domain.Employees) error
	Update(res *domain.Employees) error
	Delete() error
	ResetPasswordRequest(res *response.UserResetPasswordRequest) error
}

type EmployeeRepository interface {
	Create(ctx context.Context, employee *domain.Employees) (uint, error)
	GetByID(ctx context.Context, number uint) (*domain.Employees, error)
	Update(ctx context.Context, employee *domain.Employees) error
	NumberExist(ctx context.Context, number string) (bool, error)
	Delete(ctx context.Context, number uint) error
	GetByEmail(ctx context.Context, email string) (*domain.Employees, error)
}

type employee struct {
	outputPort   EmployeeOutputPort
	EmployeeRepo EmployeeRepository
	email        Mail
}

type EmployeeInputFactory func(outputPort EmployeeOutputPort) EmployeeInputPort

func NewEmployeeInputFactory(er EmployeeRepository, email Mail) EmployeeInputFactory {
	return func(o EmployeeOutputPort) EmployeeInputPort {
		return &employee{
			outputPort:   o,
			EmployeeRepo: er,
			email:        email,
		}
	}
}

func (e employee) Create(ctx context.Context, req *request.EmployeeCreate) error {
	number, err := e.EmployeeRepo.NumberExist(ctx, req.PhoneNumber)
	if err != nil {
		return err
	}

	if number {
		ctx.FieldError("PhoneNumber", "既に使用されています")
	}

	newUser, err := domain.NewEmployee(ctx, req)
	if err != nil {
		return err
	}

	if ctx.IsInValid() {
		return ctx.ValidationError()
	}

	id, err := e.EmployeeRepo.Create(ctx, newUser)
	if err != nil {
		return err
	}

	return e.outputPort.Create(id)
}

// make getByID code
func (e employee) GetByID(ctx context.Context, number uint) error {
	// repositoryのGetByIDを呼び出す
	res, err := e.EmployeeRepo.GetByID(ctx, number)
	if err != nil {
		return err
	}
	return e.outputPort.GetByID(res)
}

func (e employee) Update(ctx context.Context, req *request.EmployeeUpdate) error {
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

func (e employee) ResetPasswordRequest(ctx context.Context, req *request.EmployeeResetPasswordRequest) error {
	user, err := e.EmployeeRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		switch v := err.(type) {
		case *errors.Expected:
			if !v.ChangeStatus(http.StatusNotFound, http.StatusOK) {
				return err
			}
		default:
			return err
		}
	}

	var res response.UserResetPasswordRequest

	res.Duration, res.Expire, err = user.RecoveryToken.Generate()
	if err != nil {
		return err
	}

	err = ctx.Transaction(
		func(ctx context.Context) error {
			err = e.EmployeeRepo.Update(ctx, user)
			if err != nil {
				return err
			}

			err = e.email.Send(user.Email, mail_body.UserResetPasswordRequest{
				URL:   config.Env.App.URL,
				Token: user.RecoveryToken.String(),
			})
			if err != nil {
				return err
			}

			return nil
		},
	)

	if err != nil {
		return err
	}

	return e.outputPort.ResetPasswordRequest(&res)
}

func (e employee) ResetPassword(ctx context.Context, req *request.EmployeeResetPassword) error {
	//TODO implement me
	panic("implement me")
}

func (e employee) Login(ctx context.Context, req *request.EmployeeLogin) error {
	//TODO implement me
	panic("implement me")
}

func (e employee) RefreshToken(req *request.EmployeeRefreshToken) error {
	//TODO implement me
	panic("implement me")
}

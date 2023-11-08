package usecase

import (
	"attendance-management/domain"
	"attendance-management/packages/context"
	"attendance-management/resource/request"
	"time"
)

type EmploymentInputPort interface {
	Create(ctx context.Context, req *request.CreateEmployment) error
	GetByID(ctx context.Context, number uint) error
	Update(ctx context.Context, req *request.UpdateEmployment) error
	Delete(ctx context.Context, number uint) error
}

type EmploymentOutputPort interface {
	Create(id uint) error
	GetByID(res *domain.Employments) error
	Update(res *domain.Employments) error
	Delete() error
}

type EmploymentRepository interface {
	Create(ctx context.Context, employment *domain.Employments) (uint, error)
	GetByID(ctx context.Context, number uint) (*domain.Employments, error)
	Update(ctx context.Context, employment *domain.Employments) error
	NumberExist(ctx context.Context, id uint) error
	Delete(ctx context.Context, number uint) error
}

type employment struct {
	outputPort     EmploymentOutputPort
	employmentRepo EmploymentRepository
}

type EmploymentInputFactory func(outputPort EmploymentOutputPort) EmploymentInputPort

func NewEmploymentInputFactory(er EmploymentRepository) EmploymentInputFactory {
	return func(o EmploymentOutputPort) EmploymentInputPort {
		return &employment{
			outputPort:     o,
			employmentRepo: er,
		}
	}
}

func (e employment) Create(ctx context.Context, req *request.CreateEmployment) error {
	res, err := domain.NewEmployment(ctx, req)
	// req.EmploymentNumberをkeyにしてemploymentが存在するか確認
	err = e.employmentRepo.NumberExist(ctx, req.EmploymentNumber)
	if err == nil {
		return ctx.Error().BadRequest("employment already exist")
	}

	id, err := e.employmentRepo.Create(ctx, res)
	if err != nil {
		return err
	}

	return e.outputPort.Create(id)
}

func (e employment) GetByID(ctx context.Context, number uint) error {
	res, err := e.employmentRepo.GetByID(ctx, number)
	if err != nil {
		return err
	}

	return e.outputPort.GetByID(res)
}

func (e employment) Update(ctx context.Context, req *request.UpdateEmployment) error {
	employment, err := e.employmentRepo.GetByID(ctx, req.EmploymentNumber)
	if err != nil {
		return err
	}

	if req.EmployeeID != 0 {
		employment.EmployeeID = req.EmployeeID
	}
	if req.CompanyID != 0 {
		employment.CompanyID = req.CompanyID
	}
	if req.Position != "" {
		employment.Position = req.Position
	}
	if req.StartDate != (time.Time{}) {
		employment.StartDate = req.StartDate
	}

	err = e.employmentRepo.Update(ctx, employment)
	if err != nil {
		return err
	}

	return e.outputPort.Update(employment)
}

func (e employment) Delete(ctx context.Context, id uint) error {
	err := e.employmentRepo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return e.outputPort.Delete()
}

package usecase

import (
	"attendance-management/domain"
	"attendance-management/packages/context"
	"attendance-management/resource/request"
)

type SalaryTypeInputPort interface {
	Create(ctx context.Context, req *request.CreateSalaryType) error
	GetByID(ctx context.Context, number uint) error
	Update(ctx context.Context, req *request.UpdateSalaryType) error
	Delete(ctx context.Context, number uint) error
}

type SalaryTypeOutputPort interface {
	Create(id uint) error
	GetByID(res *domain.SalaryTypes) error
	Update(res *domain.SalaryTypes) error
	Delete() error
}

type SalaryTypeRepository interface {
	Create(ctx context.Context, salaryType *domain.SalaryTypes) (uint, error)
	GetByID(ctx context.Context, number uint) (*domain.SalaryTypes, error)
	Update(ctx context.Context, salaryType *domain.SalaryTypes) error
	NumberExist(ctx context.Context, number uint) error
	Delete(ctx context.Context, number uint) error
}

type salaryType struct {
	outputPort     SalaryTypeOutputPort
	salaryTypeRepo SalaryTypeRepository
}

type SalaryTypeInputFactory func(outputPort SalaryTypeOutputPort) SalaryTypeInputPort

func NewSalaryTypeInputFactory(str SalaryTypeRepository) SalaryTypeInputFactory {
	return func(o SalaryTypeOutputPort) SalaryTypeInputPort {
		return &salaryType{
			outputPort:     o,
			salaryTypeRepo: str,
		}
	}
}

func (s salaryType) Create(ctx context.Context, req *request.CreateSalaryType) error {
	newSalaryType, err := domain.NewSalaryType(ctx, req)
	// req.SalaryTypeNumberをkeyにしてsalaryTypeが存在するか確認
	err = s.salaryTypeRepo.NumberExist(ctx, req.SalaryTypeNumber)
	if err == nil {
		return ctx.Error().BadRequest("salaryType already exist")
	}

	id, err := s.salaryTypeRepo.Create(ctx, newSalaryType)
	if err != nil {
		return err
	}

	return s.outputPort.Create(id)
}

func (s salaryType) GetByID(ctx context.Context, number uint) error {
	res, err := s.salaryTypeRepo.GetByID(ctx, number)
	if err != nil {
		return err
	}

	return s.outputPort.GetByID(res)
}

func (s salaryType) Update(ctx context.Context, req *request.UpdateSalaryType) error {
	salaryType, err := s.salaryTypeRepo.GetByID(ctx, req.SalaryTypeNumber)
	if err != nil {
		return err
	}

	if req.Type != "" {
		salaryType.Type = req.Type
	}

	err = s.salaryTypeRepo.Update(ctx, salaryType)
	if err != nil {
		return err
	}

	return s.outputPort.Update(salaryType)
}

func (s salaryType) Delete(ctx context.Context, number uint) error {
	err := s.salaryTypeRepo.Delete(ctx, number)
	if err != nil {
		return err
	}

	return s.outputPort.Delete()
}

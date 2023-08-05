package usecase

import (
	"attendance-management/domain"
	"attendance-management/packages/context"
	"attendance-management/resource/request"
)

type companyInputPort interface {
	Create(ctx context.Context, req *request.CreateCompany) error
	GetByID(ctx context.Context, number uint) error
	Update(ctx context.Context, req *request.UpdateCompany) error
	Delete(ctx context.Context, number uint) error
}

type companyOutputPort interface {
	Create(id uint) error
	GetByID(res *domain.Companies) error
	Update(res *domain.Companies) error
	Delete() error
}

type companyRepository interface {
	Create(ctx context.Context, company *domain.Companies) (uint, error)
	GetByID(ctx context.Context, number uint) (*domain.Companies, error)
	Update(ctx context.Context, company *domain.Companies) error
	NumberExist(ctx context.Context, number uint) error
	Delete(ctx context.Context, number uint) error
}

type company struct {
	outputPort  companyOutputPort
	companyRepo companyRepository
}

type companyInputFactory func(outputPort companyOutputPort) companyInputPort

func NewCompanyInputFactory(cr companyRepository) companyInputFactory {
	return func(o companyOutputPort) companyInputPort {
		return &company{
			outputPort:  o,
			companyRepo: cr,
		}
	}
}

func (c company) Create(ctx context.Context, req *request.CreateCompany) error {
	newCompany, err := domain.NewCompany(ctx, req)
	// req.CompanyNumberをkeyにしてcompanyが存在するか確認
	err = c.companyRepo.NumberExist(ctx, req.CompanyNumber)
	if err == nil {
		return ctx.Error().BadRequest("company already exist")
	}

	id, err := c.companyRepo.Create(ctx, newCompany)
	if err != nil {
		return err
	}

	return c.outputPort.Create(id)
}

func (c company) GetByID(ctx context.Context, number uint) error {
	res, err := c.companyRepo.GetByID(ctx, number)
	if err != nil {
		return err
	}

	return c.outputPort.GetByID(res)
}

func (c company) Update(ctx context.Context, req *request.UpdateCompany) error {
	company, err := c.companyRepo.GetByID(ctx, req.CompanyNumber)
	if err != nil {
		return err
	}

	if req.Name != "" {
		company.Name = req.Name
	}

	err = c.companyRepo.Update(ctx, company)
	if err != nil {
		return err
	}

	return c.outputPort.Update(company)
}

func (c company) Delete(ctx context.Context, number uint) error {
	err := c.companyRepo.Delete(ctx, number)
	if err != nil {
		return err
	}

	return c.outputPort.Delete()
}

package mysql

import (
	"attendance-management/domain"
	"attendance-management/packages/context"
	"attendance-management/usecase"
)

type company struct{}

func NewCompany() usecase.CompanyRepository {
	return &company{}
}

func (c company) Create(ctx context.Context, company *domain.Companies) (uint, error) {
	db := ctx.DB()

	if err := db.Create(company).Error; err != nil {
		return 0, dbError(err)
	}
	return company.ID, nil
}

func (c company) NumberExist(ctx context.Context, number uint) error {
	db := ctx.DB()

	var company domain.Companies
	err := db.Where("number = ?", number).First(&company).Error
	if err != nil {
		return dbError(err)
	}
	return nil
}

func (c company) GetByID(ctx context.Context, number uint) (*domain.Companies, error) {
	db := ctx.DB()

	var company domain.Companies
	//ここのクエリはデータベースのカラム名と同じにすること
	err := db.Where("company_number = ?", number).First(&company).Error
	if err != nil {
		return nil, dbError(err)
	}
	return &company, nil
}

func (c company) Update(ctx context.Context, company *domain.Companies) error {
	db := ctx.DB()

	if err := db.Model(&company).Updates(company).Error; err != nil {
		return dbError(err)
	}
	return nil
}

func (c company) Delete(ctx context.Context, number uint) error {
	db := ctx.DB()

	var company domain.Companies
	res := db.Where("company_number = ?", number).Delete(&company)
	if res.Error != nil {
		return dbError(res.Error)
	}
	return nil
}

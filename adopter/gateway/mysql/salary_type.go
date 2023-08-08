package mysql

import (
	"attendance-management/domain"
	"attendance-management/packages/context"
	"attendance-management/usecase"
)

type salaryType struct{}

func NewSalaryType() usecase.SalaryTypeRepository {
	return &salaryType{}
}

func (c salaryType) Create(ctx context.Context, salaryType *domain.SalaryTypes) (uint, error) {
	db := ctx.DB()

	if err := db.Create(salaryType).Error; err != nil {
		return 0, dbError(err)
	}
	return salaryType.ID, nil
}

func (c salaryType) NumberExist(ctx context.Context, number uint) error {
	db := ctx.DB()

	var salaryType domain.SalaryTypes
	err := db.Where("number = ?", number).First(&salaryType).Error
	if err != nil {
		return dbError(err)
	}
	return nil
}

func (s salaryType) GetByID(ctx context.Context, number uint) (*domain.SalaryTypes, error) {
	db := ctx.DB()

	var salaryType domain.SalaryTypes
	//ここのクエリはデータベースのカラム名と同じにすること
	err := db.Where("salary_type_number = ?", number).First(&salaryType).Error
	if err != nil {
		return nil, dbError(err)
	}
	return &salaryType, nil
}

func (c salaryType) Update(ctx context.Context, salaryType *domain.SalaryTypes) error {
	db := ctx.DB()

	if err := db.Model(&salaryType).Updates(salaryType).Error; err != nil {
		return dbError(err)
	}
	return nil
}

func (c salaryType) Delete(ctx context.Context, number uint) error {
	db := ctx.DB()

	var salaryType domain.SalaryTypes
	res := db.Where("salary_type_number = ?", number).Delete(&salaryType)
	if res.Error != nil {
		return dbError(res.Error)
	}
	return nil
}

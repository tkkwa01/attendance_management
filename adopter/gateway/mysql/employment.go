package mysql

import (
	"attendance-management/domain"
	"attendance-management/packages/context"
	"attendance-management/usecase"
)

type employment struct{}

func NewEmployment() usecase.EmploymentRepository {
	return &employment{}
}

func (e *employment) Create(ctx context.Context, employment *domain.Employments) (uint, error) {
	db := ctx.DB()

	if err := db.Create(employment).Error; err != nil {
		return 0, dbError(err)
	}
	return employment.ID, nil
}

func (e *employment) NumberExist(ctx context.Context, number uint) error {
	db := ctx.DB()

	var employment domain.Employments
	err := db.Where("number = ?", number).First(&employment).Error
	if err != nil {
		return dbError(err)
	}
	return nil
}

func (e *employment) GetByID(ctx context.Context, number uint) (*domain.Employments, error) {
	db := ctx.DB()

	var employment domain.Employments
	//ここのクエリはデータベースのカラム名と同じにすること
	err := db.Where("employment_number = ?", number).First(&employment).Error
	if err != nil {
		return nil, dbError(err)
	}
	return &employment, nil
}

func (e *employment) Update(ctx context.Context, employment *domain.Employments) error {
	db := ctx.DB()

	if err := db.Model(&employment).Updates(employment).Error; err != nil {
		return dbError(err)
	}
	return nil
}

func (e *employment) Delete(ctx context.Context, number uint) error {
	db := ctx.DB()

	var employment domain.Employments
	res := db.Where("employment_number = ?", number).Delete(&employment)
	if res.Error != nil {
		return dbError(res.Error)
	}
	return nil
}

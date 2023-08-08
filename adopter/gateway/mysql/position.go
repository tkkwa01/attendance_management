package mysql

import (
	"attendance-management/domain"
	"attendance-management/packages/context"
	"attendance-management/usecase"
)

type position struct{}

func NewPosition() usecase.PositionRepository {
	return &position{}
}

func (p position) Create(ctx context.Context, position *domain.Positions) (uint, error) {
	db := ctx.DB()

	if err := db.Create(position).Error; err != nil {
		return 0, dbError(err)
	}
	return position.ID, nil
}

func (p position) NumberExist(ctx context.Context, number uint) error {
	db := ctx.DB()

	var position domain.Positions
	err := db.Where("number = ?", number).First(&position).Error
	if err != nil {
		return dbError(err)
	}
	return nil
}

func (p position) GetByID(ctx context.Context, number uint) (*domain.Positions, error) {
	db := ctx.DB()

	var position domain.Positions
	//ここのクエリはデータベースのカラム名と同じにすること
	err := db.Where("position_number = ?", number).First(&position).Error
	if err != nil {
		return nil, dbError(err)
	}
	return &position, nil
}

func (p position) Update(ctx context.Context, position *domain.Positions) error {
	db := ctx.DB()

	if err := db.Model(&position).Updates(position).Error; err != nil {
		return dbError(err)
	}
	return nil
}

func (p position) Delete(ctx context.Context, number uint) error {
	db := ctx.DB()

	var position domain.Positions
	res := db.Where("position_number = ?", number).Delete(&position)
	if res.Error != nil {
		return dbError(res.Error)
	}
	return nil
}

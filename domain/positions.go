package domain

import (
	"attendance-management/packages/context"
	"attendance-management/resource/request"
)

type Positions struct {
	ID             uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Type           string `json:"type" gorm:"type:varchar(255);not null"`
	PositionNumber uint   `json:"position_number" gorm:"unique"`
	Employments    []Employments
}

func NewPosition(ctx context.Context, dto *request.CreatePosition) (*Positions, error) {
	position := &Positions{
		Type:           dto.Type,
		PositionNumber: dto.PositionNumber,
	}

	if ctx.IsInValid() {
		return nil, ctx.ValidationError()
	}

	return position, nil
}

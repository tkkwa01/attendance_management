package domain

import (
	"attendance-management/packages/context"
	"attendance-management/packages/validation"
	"attendance-management/resource/request"
	"gorm.io/gorm"
)

type Positions struct {
	gorm.Model
	Type           string `json:"type"`
	PositionNumber uint   `json:"position_number" gorm:"unique"`
}

func NewPosition(ctx context.Context, req *request.CreatePosition) (*Positions, error) {
	position := &Positions{
		Type:           req.Type,
		PositionNumber: req.PositionNumber,
	}
	err := validation.Validate().Struct(position)
	if err != nil {
		return nil, err
	}
	return position, nil
}

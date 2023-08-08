package domain

import (
	"attendance-management/packages/context"
	"attendance-management/packages/validate"
	"attendance-management/resource/request"
	"gorm.io/gorm"
)

type Positions struct {
	gorm.Model
	ID             uint   `json:"id"`
	Type           string `json:"type"`
	PositionNumber uint   `json:"position_number"`
}

func NewPosition(ctx context.Context, req *request.CreatePosition) (*Positions, error) {
	position := &Positions{
		Type:           req.Type,
		PositionNumber: req.PositionNumber,
	}
	err := validate.Validate(position)
	if err != nil {
		return nil, err
	}
	return position, nil
}

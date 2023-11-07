package domain

import (
	"attendance-management/packages/context"
	"attendance-management/packages/validation"
	"attendance-management/resource/request"
)

type Positions struct {
	ID             uint          `json:"id" gorm:"primaryKey;autoIncrement"`
	Type           string        `json:"type" gorm:"type:varchar(255);not null"`
	PositionNumber uint          `json:"position_number" gorm:"unique"`
	Employments    []Employments `gorm:"foreignKey:PositionID"`
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

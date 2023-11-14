package usecase

import (
	"attendance-management/domain"
	"attendance-management/packages/context"
	"attendance-management/resource/request"
)

type PositionInputPort interface {
	Create(ctx context.Context, req *request.CreatePosition) error
	GetByID(ctx context.Context, number uint) error
	Update(ctx context.Context, req *request.UpdatePosition) error
	Delete(ctx context.Context, number uint) error
}

type PositionOutputPort interface {
	Create(id uint) error
	GetByID(res *domain.Positions) error
	Update(res *domain.Positions) error
	Delete() error
}

type PositionRepository interface {
	Create(ctx context.Context, position *domain.Positions) (uint, error)
	GetByID(ctx context.Context, number uint) (*domain.Positions, error)
	Update(ctx context.Context, position *domain.Positions) error
	NumberExist(ctx context.Context, number uint) error
	Delete(ctx context.Context, number uint) error
}

type position struct {
	outputPort   PositionOutputPort
	positionRepo PositionRepository
}

type PositionInputFactory func(outputPort PositionOutputPort) PositionInputPort

func NewPositionInputFactory(pr PositionRepository) PositionInputFactory {
	return func(o PositionOutputPort) PositionInputPort {
		return &position{
			outputPort:   o,
			positionRepo: pr,
		}
	}
}

func (p position) Create(ctx context.Context, req *request.CreatePosition) error {
	newPosition, err := domain.NewPosition(ctx, req)
	// req.PositionNumberをkeyにしてpositionが存在するか確認
	err = p.positionRepo.NumberExist(ctx, req.PositionNumber)
	if err == nil {
		return ctx.Error().BadRequest("position already exist")
	}

	id, err := p.positionRepo.Create(ctx, newPosition)
	if err != nil {
		return err
	}

	return p.outputPort.Create(id)
}

func (p position) GetByID(ctx context.Context, number uint) error {
	res, err := p.positionRepo.GetByID(ctx, number)
	if err != nil {
		return err
	}

	return p.outputPort.GetByID(res)
}

func (p position) Update(ctx context.Context, req *request.UpdatePosition) error {
	position, err := p.positionRepo.GetByID(ctx, req.PositionNumber)
	if err != nil {
		return err
	}

	if req.Type != "" {
		position.Type = req.Type
	}

	err = p.positionRepo.Update(ctx, position)
	if err != nil {
		return err
	}

	return p.outputPort.Update(position)
}

func (p position) Delete(ctx context.Context, number uint) error {
	err := p.positionRepo.Delete(ctx, number)
	if err != nil {
		return err
	}

	return p.outputPort.Delete()
}

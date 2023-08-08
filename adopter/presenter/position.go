package presenter

import (
	"attendance-management/domain"
	"attendance-management/usecase"
	"github.com/gin-gonic/gin"
)

type position struct {
	c *gin.Context
}

type PositionOutputFactory func(c *gin.Context) usecase.PositionOutputPort

func NewPositionOutputFactory() PositionOutputFactory {
	return func(c *gin.Context) usecase.PositionOutputPort {
		return &position{c: c}
	}
}

func (p *position) Create(id uint) error {
	p.c.JSON(201, gin.H{"id": id})
	return nil
}

func (p *position) GetByID(res *domain.Positions) error {
	p.c.JSON(200, res)
	return nil
}

func (p *position) Update(res *domain.Positions) error {
	p.c.JSON(200, res)
	return nil
}

func (p *position) Delete() error {
	p.c.JSON(200, "")
	return nil
}

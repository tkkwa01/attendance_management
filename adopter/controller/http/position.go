package http

import (
	"attendance-management/packages/context"
	"attendance-management/packages/http/router"
	"attendance-management/resource/request"
	"attendance-management/usecase"
	"github.com/gin-gonic/gin"
)

type position struct {
	inputFactory  usecase.PositionInputFactory
	outputFactory func(c *gin.Context) usecase.PositionOutputPort
}

func NewPosition(r *router.Router, inputFactory usecase.PositionInputFactory, outputFactory func(c *gin.Context) usecase.PositionOutputPort) {
	handler := position{
		inputFactory:  inputFactory,
		outputFactory: outputFactory,
	}

	r.Group("position", nil, func(r *router.Router) {
		r.Post("", handler.Create)
		r.Get("", handler.GetPosition)
		r.Put("", handler.Update)
		r.Delete("", handler.Delete)
	})
}

func (p position) Create(ctx context.Context, c *gin.Context) error {
	var req request.CreatePosition

	if !bind(c, &req) {
		return nil
	}

	outputPort := p.outputFactory(c)
	inputPort := p.inputFactory(outputPort)

	return inputPort.Create(ctx, &req)
}

// get position by number
func (p position) GetPosition(ctx context.Context, c *gin.Context) error {
	numberStr := c.Query("number")
	number, err := stringToUint(numberStr)
	if err != nil {
		return err
	}
	outputPort := p.outputFactory(c)
	inputPort := p.inputFactory(outputPort)

	return inputPort.GetByID(ctx, number)
}

func (p position) Update(ctx context.Context, c *gin.Context) error {
	var req request.UpdatePosition

	if !bind(c, &req) {
		return nil
	}

	outputPort := p.outputFactory(c)
	inputPort := p.inputFactory(outputPort)

	return inputPort.Update(ctx, &req)
}

func (p position) Delete(ctx context.Context, c *gin.Context) error {
	numberStr := c.Query("number")
	number, err := stringToUint(numberStr)
	if err != nil {
		return err
	}
	outputPort := p.outputFactory(c)
	inputPort := p.inputFactory(outputPort)

	return inputPort.Delete(ctx, number)
}

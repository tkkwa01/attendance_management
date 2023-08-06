package http

import (
	"attendance-management/adopter/presenter"
	"attendance-management/packages/context"
	"attendance-management/packages/http/router"
	"attendance-management/resource/request"
	"attendance-management/usecase"
	"github.com/gin-gonic/gin"
)

type employment struct {
	inputFactory  usecase.EmploymentInputFactory
	outputFactory func(c *gin.Context) usecase.EmploymentOutputPort
}

func NewEmployment(r *router.Router, inputFactory usecase.EmploymentInputFactory, outputFactory presenter.EmploymentOutputFactory) {
	handler := employment{
		inputFactory:  inputFactory,
		outputFactory: outputFactory,
	}

	r.Group("employment", nil, func(r *router.Router) {
		r.Post("", handler.Create)
		r.Get("", handler.GetEmployment)
		r.Put("", handler.Update)
		r.Delete("", handler.Delete)
	})
}

func (e employment) Create(ctx context.Context, c *gin.Context) error {
	var req request.CreateEmployment

	if !bind(c, &req) {
		return nil
	}

	outputPort := e.outputFactory(c)
	inputPort := e.inputFactory(outputPort)

	return inputPort.Create(ctx, &req)
}

func (e employment) GetEmployment(ctx context.Context, c *gin.Context) error {
	numberStr := c.Query("number")
	number, err := stringToUint(numberStr)
	if err != nil {
		return err
	}

	outputPort := e.outputFactory(c)
	inputPort := e.inputFactory(outputPort)

	return inputPort.GetByID(ctx, number)
}

func (e employment) Update(ctx context.Context, c *gin.Context) error {
	var req request.UpdateEmployment

	if !bind(c, &req) {
		return nil
	}

	outputPort := e.outputFactory(c)
	inputPort := e.inputFactory(outputPort)

	return inputPort.Update(ctx, &req)
}

func (e employment) Delete(ctx context.Context, c *gin.Context) error {
	numberStr := c.Query("number")
	number, err := stringToUint(numberStr)
	if err != nil {
		return err
	}
	outputPort := e.outputFactory(c)
	inputPort := e.inputFactory(outputPort)

	return inputPort.Delete(ctx, number)
}

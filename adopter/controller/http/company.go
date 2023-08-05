package http

import (
	"attendance-management/adopter/presenter"
	"attendance-management/packages/context"
	"attendance-management/packages/http/router"
	"attendance-management/resource/request"
	"attendance-management/usecase"
	"github.com/gin-gonic/gin"
)

type company struct {
	inputFactory  usecase.CompanyInputFactory
	outputFactory func(c *gin.Context) usecase.CompanyOutputPort
}

func NewCompany(r *router.Router, inputFactory usecase.CompanyInputFactory, outputFactory presenter.CompanyOutputFactory) {
	handler := company{
		inputFactory:  inputFactory,
		outputFactory: outputFactory,
	}

	r.Group("company", nil, func(r *router.Router) {
		r.Post("", handler.Create)
		r.Get("", handler.GetCompany)
		r.Put("", handler.Update)
		r.Delete("", handler.Delete)
	})
}

func (co company) Create(ctx context.Context, c *gin.Context) error {
	var req request.CreateCompany

	if !bind(c, &req) {
		return nil
	}

	outputPort := co.outputFactory(c)
	inputPort := co.inputFactory(outputPort)

	return inputPort.Create(ctx, &req)
}

// get company by number
func (co company) GetCompany(ctx context.Context, c *gin.Context) error {
	numberStr := c.Query("number")
	number, err := stringToUint(numberStr)
	if err != nil {
		return err
	}
	outputPort := co.outputFactory(c)
	inputPort := co.inputFactory(outputPort)

	return inputPort.GetByID(ctx, number)
}

func (co company) Update(ctx context.Context, c *gin.Context) error {
	var req request.UpdateCompany

	if !bind(c, &req) {
		return nil
	}

	outputPort := co.outputFactory(c)
	inputPort := co.inputFactory(outputPort)

	return inputPort.Update(ctx, &req)
}

func (co company) Delete(ctx context.Context, c *gin.Context) error {
	numberStr := c.Query("number")
	number, err := stringToUint(numberStr)
	if err != nil {
		return err
	}
	outputPort := co.outputFactory(c)
	inputPort := co.inputFactory(outputPort)

	return inputPort.Delete(ctx, number)
}

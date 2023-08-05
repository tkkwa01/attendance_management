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
	inputFactory  usecase.EmployeeInputFactory
	outputFactory func(c *gin.Context) usecase.EmployeeOutputPort
}

func NewCompany(r *router.Router, inputFactory usecase.EmployeeInputFactory, outputFactory presenter.EmployeeOutputFactory) {
	handler := company{
		inputFactory:  inputFactory,
		outputFactory: outputFactory,
	}

	r.Group("company", nil, func(r *router.Router) {
		r.Post("", handler.Create)
		r.Get("", handler.GetEmployee)
		r.Put("", handler.Update)
		r.Delete("", handler.Delete)
	})
}

func (co company) Create(ctx context.Context, c *gin.Context) error {
	var req request.CreateCompany

	if !bind(co, &req) {
		return nil
	}

	outputPort := co.outputFactory(c)
	inputPort := co.inputFactory(outputPort)

	return inputPort.Create(ctx, &req)
}

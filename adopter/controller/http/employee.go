package http

import (
	"attendance-management/adopter/controller/presenter"
	"attendance-management/packages/context"
	"attendance-management/packages/http/router"
	"attendance-management/resource/request"
	"attendance-management/usecase"
	"github.com/gin-gonic/gin"
)

type employee struct {
	inputFactory  usecase.EmployeeInputFactory
	outputFactory func(c *gin.Context) usecase.EmployeeOutputPort
}

func NewEmployee(r *router.Router, inputFactory usecase.EmployeeInputFactory, outputFactory presenter.EmployeeOutputFactory) {
	handler := employee{
		inputFactory:  inputFactory,
		outputFactory: outputFactory,
	}

	r.Group("employee", nil, func(r *router.Router) {
		r.Post("", handler.Create)
	})
}

func (e employee) Create(ctx context.Context, c *gin.Context) error {
	var req request.CreateEmployee

	if !bind(c, &req) {
		return nil
	}

	outputPort := e.outputFactory(c)
	inputPort := e.inputFactory(outputPort)

	return inputPort.Create(ctx, &req)
}

// get employee by number
func (e employee) GetEmployee(ctx context.Context, c *gin.Context) error {
	number := c.Query("number")
	outputPort := e.outputFactory(c)
	inputPort := e.inputFactory(outputPort)

	return inputPort.GetByID(ctx, number)
}

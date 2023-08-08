package http

import (
	"attendance-management/packages/context"
	"attendance-management/packages/http/router"
	"attendance-management/resource/request"
	"attendance-management/usecase"
	"github.com/gin-gonic/gin"
)

type salaryType struct {
	inputFactory  usecase.SalaryTypeInputFactory
	outputFactory func(c *gin.Context) usecase.SalaryTypeOutputPort
}

func NewSalaryType(r *router.Router, inputFactory usecase.SalaryTypeInputFactory, outputFactory func(c *gin.Context) usecase.SalaryTypeOutputPort) {
	handler := salaryType{
		inputFactory:  inputFactory,
		outputFactory: outputFactory,
	}

	r.Group("salary_type", nil, func(r *router.Router) {
		r.Post("", handler.Create)
		r.Get("", handler.GetSalaryType)
		r.Put("", handler.Update)
		r.Delete("", handler.Delete)
	})
}

func (s salaryType) Create(ctx context.Context, c *gin.Context) error {
	var req request.CreateSalaryType

	if !bind(c, &req) {
		return nil
	}

	outputPort := s.outputFactory(c)
	inputPort := s.inputFactory(outputPort)

	return inputPort.Create(ctx, &req)
}

// get salary_type by number
func (s salaryType) GetSalaryType(ctx context.Context, c *gin.Context) error {
	numberStr := c.Query("number")
	number, err := stringToUint(numberStr)
	if err != nil {
		return err
	}
	outputPort := s.outputFactory(c)
	inputPort := s.inputFactory(outputPort)

	return inputPort.GetByID(ctx, number)
}

func (s salaryType) Update(ctx context.Context, c *gin.Context) error {
	var req request.UpdateSalaryType

	if !bind(c, &req) {
		return nil
	}

	outputPort := s.outputFactory(c)
	inputPort := s.inputFactory(outputPort)

	return inputPort.Update(ctx, &req)
}

func (s salaryType) Delete(ctx context.Context, c *gin.Context) error {
	numberStr := c.Query("number")
	number, err := stringToUint(numberStr)
	if err != nil {
		return err
	}
	outputPort := s.outputFactory(c)
	inputPort := s.inputFactory(outputPort)

	return inputPort.Delete(ctx, number)
}

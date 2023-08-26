package http

import (
	"attendance-management/packages/context"
	"attendance-management/packages/http/router"
	"attendance-management/resource/request"
	"attendance-management/usecase"
	"github.com/gin-gonic/gin"
)

type attendance struct {
	inputFactory  usecase.AttendanceInputFactory
	outputFactory func(c *gin.Context) usecase.AttendanceOutputPort
}

func NewAttendance(r *router.Router, inputFactory usecase.AttendanceInputFactory, outputFactory func(c *gin.Context) usecase.AttendanceOutputPort) {
	handler := attendance{
		inputFactory:  inputFactory,
		outputFactory: outputFactory,
	}

	r.Group("attendance", nil, func(r *router.Router) {
		r.Post("in", handler.Create)
		r.Put("out", handler.Update)
		r.Get("", handler.GetAttendance)
		r.Put("", handler.Update)
		r.Delete("", handler.Delete)
	})
}

func (a attendance) Create(ctx context.Context, c *gin.Context) error {
	var req request.CreateAttendance

	if !bind(c, &req) {
		return nil
	}

	outputPort := a.outputFactory(c)
	inputPort := a.inputFactory(outputPort)

	return inputPort.Create(ctx, &req)
}

func (a attendance) GetAttendance(ctx context.Context, c *gin.Context) error {
	numberStr := c.Query("number")
	number, err := stringToUint(numberStr)
	if err != nil {
		return err
	}

	outputPort := a.outputFactory(c)
	inputPort := a.inputFactory(outputPort)

	return inputPort.GetByID(ctx, number)
}

func (a attendance) Update(ctx context.Context, c *gin.Context) error {
	var req request.UpdateAttendance

	if !bind(c, &req) {
		return nil
	}

	outputPort := a.outputFactory(c)
	inputPort := a.inputFactory(outputPort)

	return inputPort.Update(ctx, &req)
}

func (a attendance) Delete(ctx context.Context, c *gin.Context) error {
	numberStr := c.Query("number")
	number, err := stringToUint(numberStr)
	if err != nil {
		return err
	}

	outputPort := a.outputFactory(c)
	inputPort := a.inputFactory(outputPort)

	return inputPort.Delete(ctx, number)
}

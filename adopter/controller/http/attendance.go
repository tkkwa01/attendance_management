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

	r.Group("attendances", nil, func(r *router.Router) {
		r.Post("check-in", handler.CheckIn)
		r.Post("check-out", handler.CheckOut)
		r.Get("", handler.GetAttendance)
		r.Put("", handler.Update)
		r.Delete("", handler.Delete)
	})
}

func (a attendance) CheckIn(ctx context.Context, c *gin.Context) error {
	var req request.CreateAttendance
	number := req.AttendanceNumber

	if !bind(c, &req) {
		return nil
	}

	outputPort := a.outputFactory(c)
	inputPort := a.inputFactory(outputPort)

	return inputPort.CheckIn(ctx, &req, number)
}

func (a attendance) CheckOut(ctx context.Context, c *gin.Context) error {
	var req request.CreateAttendance
	number := req.AttendanceNumber

	if !bind(c, &req) {
		return nil
	}

	outputPort := a.outputFactory(c)
	inputPort := a.inputFactory(outputPort)

	return inputPort.CheckOut(ctx, number)
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

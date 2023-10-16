package http

import (
	"attendance-management/packages/context"
	"attendance-management/packages/http/router"
	"attendance-management/resource/request"
	"attendance-management/usecase"
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
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
		r.Post("check-in/:employee_number", handler.CheckIn)
		r.Put("check-out/:attendance_number", handler.CheckOut)
		r.Get("", handler.GetAttendance)
		r.Get("all", handler.GetAll)
		r.Put(":employee_number", handler.Update)
		r.Delete(":employee_number", handler.Delete)
	})
}

func (a attendance) CheckIn(ctx context.Context, c *gin.Context) error {
	var req request.CreateAttendance

	// IDを文字列として取得
	employeeNumberStr := c.Param("employee_number")
	if employeeNumberStr == "" {
		return errors.New("employee_number parameter is missing")
	}

	// 文字列をuintに変換
	employeeNumber, err := strconv.ParseUint(employeeNumberStr, 10, 64)
	if err != nil {
		return errors.New("invalid employee_number parameter")
	}
	req.EmployeeNumber = uint(employeeNumber)

	// リクエストボディをバインド
	if err := c.Bind(&req); err != nil {
		return err
	}

	number := req.AttendanceNumber

	outputPort := a.outputFactory(c)
	inputPort := a.inputFactory(outputPort)

	return inputPort.CheckIn(ctx, &req, number)
}

func (a attendance) CheckOut(ctx context.Context, c *gin.Context) error {
	var req request.CheckOutAttendance

	// IDを文字列として取得
	employeeNumberStr := c.Param("attendance_number")
	if employeeNumberStr == "" {
		return errors.New("attendance_number parameter is missing")
	}
	// 文字列をuintに変換
	attendanceNumber, err := strconv.ParseUint(employeeNumberStr, 10, 64)
	if err != nil {
		return errors.New("invalid attendance_number parameter")
	}
	req.AttendanceNumber = uint(attendanceNumber)
	// リクエストボディをバインド
	if err := c.Bind(&req); err != nil {
		return err
	}

	number := req.AttendanceNumber

	outputPort := a.outputFactory(c)
	inputPort := a.inputFactory(outputPort)

	return inputPort.CheckOut(ctx, number)
}

func (a attendance) GetAttendance(ctx context.Context, c *gin.Context) error {
	numberStr := c.Query("attendance_number")
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

func (a attendance) GetAll(ctx context.Context, c *gin.Context) error {
	outputPort := a.outputFactory(c)
	inputPort := a.inputFactory(outputPort)

	return inputPort.GetAll(ctx)
}

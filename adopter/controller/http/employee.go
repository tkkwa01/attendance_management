package http

import (
	"attendance-management/adopter/presenter"
	"attendance-management/config"
	"attendance-management/packages/context"
	"attendance-management/packages/http/middleware"
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
		r.Post("login", handler.Login)
		r.Post("refresh-token", handler.RefreshToken)
		r.Patch("reset-password-request", handler.ResetPasswordRequest)
		r.Patch("reset-password", handler.ResetPassword)
	})

	r.Group("", []gin.HandlerFunc{middleware.Auth(true, config.UserRealm, true)}, func(r *router.Router) {
		r.Group("employee", nil, func(r *router.Router) {
			r.Get("me", handler.GetMe)
		})
	})
}

func (e employee) Create(ctx context.Context, c *gin.Context) error {
	var req request.EmployeeCreate

	if !bind(c, &req) {
		return nil
	}

	outputPort := e.outputFactory(c)
	inputPort := e.inputFactory(outputPort)

	return inputPort.Create(ctx, &req)
}

func (e employee) GetMe(ctx context.Context, c *gin.Context) error {
	outputPort := e.outputFactory(c)
	inputPort := e.inputFactory(outputPort)

	return inputPort.GetByID(ctx, ctx.UID())
}

func (e employee) Update(ctx context.Context, c *gin.Context) error {
	var req request.EmployeeUpdate

	if !bind(c, &req) {
		return nil
	}

	outputPort := e.outputFactory(c)
	inputPort := e.inputFactory(outputPort)

	return inputPort.Update(ctx, &req)
}

func (e employee) Delete(ctx context.Context, c *gin.Context) error {
	numberStr := c.Query("number")
	number, err := stringToUint(numberStr)
	if err != nil {
		return err
	}
	outputPort := e.outputFactory(c)
	inputPort := e.inputFactory(outputPort)

	return inputPort.Delete(ctx, number)
}

func (e employee) ResetPasswordRequest(ctx context.Context, c *gin.Context) error {
	var req request.EmployeeResetPasswordRequest

	if !bind(c, &req) {
		return nil
	}

	outputPort := e.outputFactory(c)
	inputPort := e.inputFactory(outputPort)

	return inputPort.ResetPasswordRequest(ctx, &req)
}

func (e employee) ResetPassword(ctx context.Context, c *gin.Context) error {
	var req request.EmployeeResetPassword

	if !bind(c, &req) {
		return nil
	}

	outputPort := e.outputFactory(c)
	inputPort := e.inputFactory(outputPort)

	return inputPort.ResetPassword(ctx, &req)
}

func (e employee) Login(ctx context.Context, c *gin.Context) error {
	var req request.EmployeeLogin

	if !bind(c, &req) {
		return nil
	}

	outputPort := e.outputFactory(c)
	inputPort := e.inputFactory(outputPort)

	return inputPort.Login(ctx, &req)
}

func (e employee) RefreshToken(_ context.Context, c *gin.Context) error {
	var req request.EmployeeRefreshToken

	if !bind(c, &req) {
		return nil
	}

	outputPort := e.outputFactory(c)
	inputPort := e.inputFactory(outputPort)

	return inputPort.RefreshToken(&req)
}

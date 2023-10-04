package presenter

import (
	"attendance-management/config"
	"attendance-management/domain"
	"attendance-management/resource/response"
	"attendance-management/usecase"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type employee struct {
	c *gin.Context
}

type EmployeeOutputFactory func(c *gin.Context) usecase.EmployeeOutputPort

func NewEmployeeOutputFactory() EmployeeOutputFactory {
	return func(c *gin.Context) usecase.EmployeeOutputPort {
		return &employee{c: c}
	}
}

func (e *employee) Create(id uint) error {
	e.c.JSON(http.StatusCreated, gin.H{"id": id})
	return nil
}

func (e *employee) GetByID(res *domain.Employees) error {
	e.c.JSON(http.StatusOK, res)
	return nil
}

func (e *employee) Update(res *domain.Employees) error {
	e.c.JSON(http.StatusOK, res)
	return nil
}

func (e *employee) Delete() error {
	e.c.JSON(http.StatusOK, "")
	return nil
}

func (e *employee) ResetPasswordRequest(res *response.UserResetPasswordRequest) error {
	e.c.JSON(http.StatusOK, res)
	return nil
}

func (e *employee) ResetPassword() error {
	e.c.Status(http.StatusOK)
	return nil
}

func (e *employee) Login(isSession bool, res *response.UserLogin) error {
	if res == nil {
		e.c.Status(http.StatusUnauthorized)
		return nil
	}

	if isSession {
		session := sessions.DefaultMany(e.c, config.UserRealm)
		session.Set("token", res.Token)
		session.Set("refresh_token", res.RefreshToken)
		if err := session.Save(); err != nil {
			return err
		}
		e.c.Status(http.StatusOK)
	} else {
		e.c.JSON(http.StatusOK, res)
	}

	return nil
}

func (e *employee) RefreshToken(isSession bool, res *response.UserLogin) error {
	if res == nil {
		e.c.Status(http.StatusUnauthorized)
		return nil
	}

	if isSession {
		session := sessions.DefaultMany(e.c, config.UserRealm)
		session.Set("token", res.Token)
		session.Set("refresh_token", res.RefreshToken)
		if err := session.Save(); err != nil {
			return err
		}
		e.c.Status(http.StatusOK)
	} else {
		e.c.JSON(http.StatusOK, res)
	}

	return nil
}

package presenter

import (
	"attendance-management/usecase"
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

func (e *employee) GetByID() error {
	//TODO implement me
	panic("implement me")
}

func (e *employee) Update() error {
	//TODO implement me
	panic("implement me")
}

func (e *employee) Delete() error {
	//TODO implement me
	panic("implement me")
}

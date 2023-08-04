package presenter

import (
	"attendance-management/domain"
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

func (e *employee) GetByID(res *domain.Employees) error {
	e.c.JSON(http.StatusOK, res)
	return nil
}

func (e *employee) Update(res *domain.Employees) error {
	e.c.JSON(http.StatusOK, res)
	return nil
}

func (e *employee) Delete() error {
	e.c.JSON(http.StatusOK, gin.H{"message": "delete employee"})
	return nil
}

package presenter

import (
	"attendance-management/domain"
	"attendance-management/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type employment struct {
	c *gin.Context
}

type EmploymentOutputFactory func(c *gin.Context) usecase.EmploymentOutputPort

func NewEmploymentOutputFactory() EmploymentOutputFactory {
	return func(c *gin.Context) usecase.EmploymentOutputPort {
		return &employment{c: c}
	}
}

func (e *employment) Create(id uint) error {
	e.c.JSON(http.StatusCreated, gin.H{"id": id})
	return nil
}

func (e *employment) GetByID(res *domain.Employments) error {
	e.c.JSON(http.StatusOK, res)
	return nil
}

func (e *employment) Update(res *domain.Employments) error {
	e.c.JSON(http.StatusOK, res)
	return nil
}

func (e *employment) Delete() error {
	e.c.JSON(http.StatusOK, "")
	return nil
}

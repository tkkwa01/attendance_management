package presenter

import (
	"attendance-management/domain"
	"attendance-management/usecase"
	"github.com/gin-gonic/gin"
)

type salaryType struct {
	c *gin.Context
}

type SalaryTypeOutputFactory func(c *gin.Context) usecase.SalaryTypeOutputPort

func NewSalaryTypeOutputFactory() SalaryTypeOutputFactory {
	return func(c *gin.Context) usecase.SalaryTypeOutputPort {
		return &salaryType{c: c}
	}
}

func (s *salaryType) Create(id uint) error {
	s.c.JSON(201, gin.H{"id": id})
	return nil
}

func (s *salaryType) GetByID(res *domain.SalaryTypes) error {
	s.c.JSON(200, res)
	return nil
}

func (s *salaryType) Update(res *domain.SalaryTypes) error {
	s.c.JSON(200, res)
	return nil
}

func (s *salaryType) Delete() error {
	s.c.JSON(200, "")
	return nil
}

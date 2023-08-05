package presenter

import (
	"attendance-management/domain"
	"attendance-management/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type company struct {
	c *gin.Context
}

type CompanyOutputFactory func(c *gin.Context) usecase.CompanyOutputPort

func NewCompanyOutputFactory() CompanyOutputFactory {
	return func(c *gin.Context) usecase.CompanyOutputPort {
		return &company{c: c}
	}
}

func (c *company) Create(id uint) error {
	c.c.JSON(201, gin.H{"id": id})
	return nil
}

func (c *company) GetByID(res *domain.Companies) error {
	c.c.JSON(http.StatusOK, res)
	return nil
}

func (c *company) Update(res *domain.Companies) error {
	c.c.JSON(http.StatusOK, res)
	return nil
}

func (c *company) Delete() error {
	c.c.JSON(http.StatusOK, "")
	return nil
}

package presenter

import (
	"attendance-management/domain"
	"attendance-management/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type attendance struct {
	c *gin.Context
}

type AttendanceOutputFactory func(c *gin.Context) usecase.AttendanceOutputPort

func NewAttendanceOutputFactory() AttendanceOutputFactory {
	return func(c *gin.Context) usecase.AttendanceOutputPort {
		return &attendance{c: c}
	}
}

func (a *attendance) CheckIn(id uint) error {
	a.c.JSON(201, gin.H{"id": id})
	return nil
}

func (a *attendance) CheckOut(res *domain.Attendance) error {
	a.c.JSON(200, res)
	return nil
}

func (a *attendance) Create(id uint) error {
	a.c.JSON(201, gin.H{"id": id})
	return nil
}

func (a *attendance) GetByID(res *domain.Attendance) error {
	a.c.JSON(200, res)
	return nil
}

func (a *attendance) Update(res *domain.Attendance) error {
	a.c.JSON(200, res)
	return nil
}

func (a *attendance) Delete() error {
	a.c.JSON(200, "")
	return nil
}

func (a *attendance) GetAll(res []*domain.Attendance) error {
	a.c.JSON(http.StatusOK, res)
	return nil
}

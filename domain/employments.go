package domain

import (
	"attendance-management/packages/context"
	"attendance-management/resource/request"
	"time"
)

type Employments struct {
	ID               uint         `json:"id" gorm:"primaryKey;autoIncrement"`
	EmployeeID       uint         `json:"employee_id" gorm:"not null"`
	CompanyID        uint         `json:"company_id" gorm:"not null"`
	Position         string       `json:"position_id" gorm:"type:varchar(255);not null"`
	StartDate        time.Time    `json:"start_date" gorm:"type:date;not null"`
	EndDate          *time.Time   `json:"end_date" gorm:"type:date"`
	SalaryTypeID     uint         `json:"salary_type_id" gorm:"not null"`
	EmploymentNumber uint         `json:"employment_number" gorm:"unique"`
	Attendance       []Attendance `json:"attendance" gorm:"foreignKey:EmploymentID"`
	Salaries         []Salaries   `json:"salaries" gorm:"foreignKey:EmploymentID"`
}

func NewEmployment(ctx context.Context, dto *request.CreateEmployment) (*Employments, error) {
	var employments = &Employments{
		EmployeeID:       dto.EmployeeID,
		CompanyID:        dto.CompanyID,
		Position:         dto.Position,
		StartDate:        dto.StartDate,
		SalaryTypeID:     dto.SalaryTypeID,
		EmploymentNumber: dto.EmploymentNumber,
	}

	if ctx.IsInValid() {
		return nil, ctx.ValidationError()
	}

	return employments, nil
}

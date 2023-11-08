package domain

import (
	"time"
)

type Salaries struct {
	ID                 uint          `json:"id" gorm:"primaryKey;autoIncrement"`
	EmploymentID       uint          `json:"employment_id" gorm:"not null"`
	MonthYear          time.Time     `json:"month_year" gorm:"type:date;not null"`
	MonthlyTotalSalary int           `json:"monthly_total_salary"`
	HoursWorked        int           `json:"hours_worked"`
	Employments        []Employments `gorm:"many2many:employments"`
}

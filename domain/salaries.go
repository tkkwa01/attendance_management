package domain

import (
	"gorm.io/gorm"
	"time"
)

type Salaries struct {
	gorm.Model
	ID                 uint      `json:"id"`
	EmploymentID       uint      `json:"employment_id"`
	MonthYear          time.Time `json:"month_year"`
	MonthlyTotalSalary int       `json:"monthly_total_salary"`
	HoursWorked        int       `json:"hours_worked"`
}

package domain

import (
	"time"
)

type Salaries struct {
	ID                 uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	EmploymentID       uint      `json:"employment_id" gorm:"not null;foreignKey"`
	MonthYear          time.Time `json:"month_year" gorm:"type:date;not null"`
	MonthlyTotalSalary int       `json:"monthly_total_salary"`
	HoursWorked        int       `json:"hours_worked"`
}

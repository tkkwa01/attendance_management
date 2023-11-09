package request

import "time"

type CreateEmployment struct {
	EmployeeID       uint       `json:"employee_id"`
	CompanyID        uint       `json:"company_id"`
	PositionID       uint       `json:"position_id"`
	Position         string     `json:"position"`
	StartDate        time.Time  `json:"start_date"`
	EndDate          *time.Time `json:"end_date"`
	SalaryTypeID     uint       `json:"salary_type_id"`
	EmploymentNumber uint       `json:"employment_number"`
}

type UpdateEmployment struct {
	ID               uint       `json:"id"`
	EmployeeID       uint       `json:"employee_id"`
	CompanyID        uint       `json:"company_id"`
	PositionID       uint       `json:"position_id"`
	StartDate        time.Time  `json:"start_date"`
	EndDate          *time.Time `json:"end_date"`
	SalaryTypeID     uint       `json:"salary_type_id"`
	EmploymentNumber uint       `json:"employment_number"`
}

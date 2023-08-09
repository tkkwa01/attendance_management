package domain

import (
	"attendance-management/packages/context"
	"attendance-management/packages/validate"
	"attendance-management/resource/request"
	"gorm.io/gorm"
	"time"
)

type Employments struct {
	gorm.Model
	ID               uint      `json:"id"`
	EmployeeID       uint      `json:"employee_id"`
	CompanyID        uint      `json:"company_id"`
	PositionID       uint      `json:"position_id"`
	StartDate        time.Time `json:"start_date"`
	EndDate          time.Time `json:"end_date"`
	SalaryTypeID     uint      `json:"salary_type_id"`
	EmploymentNumber uint      `json:"employment_number" gorm:"unique"`
	//Employees        Employees   `gorm:"foreignKey:EmployeeID"`
	//Companies        Companies   `gorm:"foreignKey:CompanyID"`
	//Positions        Positions   `gorm:"foreignKey:PositionID"`
	//SalaryTypes      SalaryTypes `gorm:"foreignKey:SalaryTypeID"`
}

func NewEmployment(ctx context.Context, req *request.CreateEmployment) (*Employments, error) {
	employments := &Employments{
		ID:               req.ID,
		EmployeeID:       req.EmployeeID,
		CompanyID:        req.CompanyID,
		PositionID:       req.PositionID,
		StartDate:        req.StartDate,
		EndDate:          req.EndDate,
		SalaryTypeID:     req.SalaryTypeID,
		EmploymentNumber: req.EmploymentNumber,
	}
	err := validate.Validate(employments)
	if err != nil {
		return nil, err
	}
	return employments, nil
}

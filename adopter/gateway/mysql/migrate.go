package mysql

import (
	"attendance-management/domain"
	"attendance-management/driver"
)

func init() {
	err := driver.GetRDB().AutoMigrate(
		&domain.Attendance{},
		&domain.Companies{},
		&domain.Employees{},
		&domain.Employments{},
		&domain.Positions{},
		&domain.Salaries{},
		&domain.SalaryTypes{},
	)
	if err != nil {
		panic(err)
	}
}

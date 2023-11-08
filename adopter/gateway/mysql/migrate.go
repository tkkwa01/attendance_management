package mysql

import (
	"attendance-management/domain"
	"attendance-management/driver"
)

func init() {
	err := driver.GetRDB().AutoMigrate(
		&domain.Positions{},
		&domain.SalaryTypes{},
		&domain.Companies{},
		&domain.Employees{},
		&domain.Employments{},
		&domain.Attendance{},
		&domain.Salaries{},
	)
	if err != nil {
		panic(err)
	}
}

package mysql

import (
	"attendance-management/domain"
	"attendance-management/packages/context"
	"attendance-management/usecase"
)

type employee struct{}

func NewEmployee() usecase.EmployeeRepository {
	return &employee{}
}

func (e employee) Create(ctx context.Context, employee *domain.Employees) (uint, error) {
	db := ctx.DB()

	if err := db.Create(employee).Error; err != nil {
		return 0, dbError(err)
	}
	return employee.ID, nil
}

func (e employee) NumberExist(ctx context.Context, number uint) error {
	db := ctx.DB()

	var employee domain.Employees
	err := db.Where("number = ?", number).First(&employee).Error
	if err != nil {
		return dbError(err)
	}
	return nil
}
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

func (e employee) NumberExist(ctx context.Context, number string) (bool, error) {
	db := ctx.DB()

	var count int64
	if err := db.Model(&domain.Employees{}).Where(&domain.Employees{PhoneNumber: number}).Count(&count).Error; err != nil {
		return false, dbError(err)
	}
	return count > 0, nil
}

func (e employee) GetByID(ctx context.Context, number uint) (*domain.Employees, error) {
	db := ctx.DB()

	var employee domain.Employees
	//ここのクエリはデータベースのカラム名と同じにすること
	err := db.Where("employee_number = ?", number).First(&employee).Error
	if err != nil {
		return nil, dbError(err)
	}
	return &employee, nil
}

func (e employee) Update(ctx context.Context, employee *domain.Employees) error {
	db := ctx.DB()

	if err := db.Model(&employee).Updates(employee).Error; err != nil {
		return dbError(err)
	}
	return nil
}

func (e employee) Delete(ctx context.Context, number uint) error {
	db := ctx.DB()

	var employee domain.Employees
	res := db.Where("employee_number = ?", number).Delete(&employee)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (e employee) GetByEmail(ctx context.Context, email string) (*domain.Employees, error) {
	db := ctx.DB()

	var dest domain.Employees
	err := db.Where(&domain.Employees{Email: email}).First(&dest).Error
	if err != nil {
		return nil, dbError(err)
	}
	return &dest, nil
}

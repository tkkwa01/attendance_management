package mysql

import (
	"attendance-management/domain"
	"attendance-management/domain/vobj"
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

func (e employee) GetByID(ctx context.Context, id uint) (*domain.Employees, error) {
	db := ctx.DB()

	var user domain.Employees
	err := db.Where(&domain.Employees{ID: id}).First(&user).Error
	if err != nil {
		return nil, dbError(err)
	}
	return &user, nil
}

func (e employee) Update(ctx context.Context, employee *domain.Employees) error {
	db := ctx.DB()

	if err := db.Model(&employee).Updates(employee).Error; err != nil {
		return dbError(err)
	}
	return nil
}

func (e employee) Delete(ctx context.Context, id uint) error {
	db := ctx.DB()

	var employee domain.Employees
	res := db.Where("id = ?", id).Delete(&employee)
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

func (e employee) EmailExists(ctx context.Context, email string) (bool, error) {
	db := ctx.DB()

	var count int64
	if err := db.Model(&domain.Employees{}).Where(&domain.Employees{Email: email}).Count(&count).Error; err != nil {
		return false, dbError(err)
	}
	return count > 0, nil
}

func (e employee) GetByRecoveryToken(ctx context.Context, recoveryToken string) (*domain.Employees, error) {
	db := ctx.DB()

	var dest domain.Employees
	err := db.Where(&domain.Employees{RecoveryToken: vobj.NewRecoveryToken(recoveryToken)}).First(&dest).Error
	if err != nil {
		return nil, dbError(err)
	}
	return &dest, nil
}

func (e employee) GetAll(ctx context.Context) ([]*domain.Employees, error) {
	db := ctx.DB()

	// 従業員を保存するスライスを定義
	var employees []*domain.Employees

	// 全ての従業員をデータベースから取得
	err := db.Find(&employees).Error
	if err != nil {
		return nil, dbError(err)
	}

	return employees, nil
}

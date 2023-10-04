package domain

import (
	"attendance-management/domain/vobj"
	"attendance-management/packages/context"
	"attendance-management/resource/request"
	"gorm.io/gorm"
)

type Employees struct {
	gorm.Model
	ID             uint                `json:"id"`
	Name           string              `json:"name"`
	PhoneNumber    string              `json:"phone_number"`
	Email          string              `json:"email" gorm:"index;unique"`
	EmployeeNumber uint                `json:"employee_number" gorm:"unique"`
	RecoveryToken  *vobj.RecoveryToken `json:"-" gorm:"index"`
	Password       vobj.Password       `json:"-"`
}

func NewEmployee(ctx context.Context, dto *request.EmployeeCreate) (*Employees, error) {
	var user = Employees{
		Name:           dto.Name,
		PhoneNumber:    dto.PhoneNumber,
		EmployeeNumber: dto.EmployeeNumber,
		RecoveryToken:  vobj.NewRecoveryToken(""),
	}

	ctx.Validate(user)

	password, err := vobj.NewPassword(ctx, dto.Password, dto.PasswordConfirm)
	if err != nil {
		return nil, err
	}

	user.Password = *password

	return &user, nil
}

func (e *Employees) ResetPassword(ctx context.Context, dto *request.EmployeeResetPassword) error {
	if !e.RecoveryToken.IsValid() {
		ctx.FieldError("RecoveryToken", "リカバリートークンが無効です")
		return nil
	}

	password, err := vobj.NewPassword(ctx, dto.Password, dto.PasswordConfirm)
	if err != nil {
		return err
	}

	e.Password = *password

	e.RecoveryToken.Clear()
	return nil
}

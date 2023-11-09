package domain

import (
	"attendance-management/domain/vobj"
	"attendance-management/packages/context"
	"attendance-management/resource/request"
)

type Employees struct {
	ID            uint                `json:"id" gorm:"primaryKey;autoIncrement"`
	Name          string              `json:"name" gorm:"type:varchar(255);not null"`
	PhoneNumber   string              `json:"phone_number" gorm:"type:varchar(255);not null"`
	Email         string              `json:"email" gorm:"index;unique"`
	RecoveryToken *vobj.RecoveryToken `json:"-" gorm:"index"`
	Password      vobj.Password       `json:"-"`
	Employments   []Employments       `json:"employments" gorm:"foreignKey:EmployeeID"`
}

func NewEmployee(ctx context.Context, dto *request.EmployeeCreate) (*Employees, error) {
	var user = Employees{
		Name:          dto.Name,
		PhoneNumber:   dto.PhoneNumber,
		Email:         dto.Email,
		RecoveryToken: vobj.NewRecoveryToken(""),
	}

	if ctx.IsInValid() {
		return nil, ctx.ValidationError()
	}

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

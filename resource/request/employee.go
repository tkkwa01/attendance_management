package request

type EmployeeCreate struct {
	Name            string `json:"name"`
	PhoneNumber     string `json:"phone_number"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}

type EmployeeUpdate struct {
	Name           string `json:"name"`
	EmployeeNumber uint   `json:"employee_number"`
	PhoneNumber    string `json:"phone_number"`
	Email          string `json:"email"`
}

type EmployeeLogin struct {
	Session  bool   `json:"session"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type EmployeeRefreshToken struct {
	Session      bool   `json:"session"`
	RefreshToken string `json:"refresh_token"`
}

type EmployeeResetPasswordRequest struct {
	Email string `json:"email"`
}

type EmployeeResetPassword struct {
	RecoveryToken   string `json:"recovery_token"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}

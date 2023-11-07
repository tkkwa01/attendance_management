package request

type CreateSalaryType struct {
	Type             string `json:"type" validate:"required"`
	SalaryTypeNumber uint   `json:"salary_type_number"`
}

type UpdateSalaryType struct {
	Type             string `json:"type"`
	SalaryTypeNumber uint   `json:"salary_type_number"`
}

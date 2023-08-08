package request

type CreateSalaryType struct {
	Type             string `json:"type"`
	SalaryTypeNumber uint   `json:"salaly_type_number"`
}

type UpdateSalaryType struct {
	Type             string `json:"type"`
	SalaryTypeNumber uint   `json:"salaly_type_number"`
}

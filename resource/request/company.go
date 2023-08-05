package request

type CreateCompany struct {
	Name          string `json:"name"`
	CompanyNumber uint   `json:"company_number"`
}

type UpdateCompany struct {
	Name          string `json:"name"`
	CompanyNumber uint   `json:"company_number"`
}

package request

type CreateEmployee struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

type UpdateEmployee struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

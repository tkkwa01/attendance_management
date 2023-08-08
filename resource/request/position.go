package request

type CreatePosition struct {
	Type           string `json:"type"`
	PositionNumber uint   `json:"position_number"`
}

type UpdatePosition struct {
	Type           string `json:"type"`
	PositionNumber uint   `json:"position_number"`
}

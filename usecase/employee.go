package usecase

type EmployeeInputPort interface {
	Create() error
	GetByID() error
	Update() error
	Delete() error
}

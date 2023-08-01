package errors

import "fmt"

type Error struct {
	Message string
	Fields  map[string]string
}

func New() *Error {
	return &Error{
		Fields: make(map[string]string),
	}
}

func (e *Error) AddError(field string, message string) {
	e.Fields[field] = message
}

func (e *Error) Error() string {
	if e.Message != "" {
		return e.Message
	}
	return fmt.Sprintf("Validation errors: %v", e.Fields)
}

func (e *Error) BadRequest(s string) error {
	e.Message = s
	return e
}

func NotFound() *Error {
	return &Error{Message: "Resource not found"}
}

func NewUnexpected(err error) *Error {
	return &Error{Message: fmt.Sprintf("Unexpected error: %v", err)}
}

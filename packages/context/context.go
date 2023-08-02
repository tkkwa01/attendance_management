package context

import (
	"attendance-management/packages/errors"
	"attendance-management/packages/validate"
	"github.com/go-playground/validator/v10"
)

type context interface {
	Error() *errors.Error
	Validate(i interface{}) error
}

type ctx struct {
	err *errors.Error
}

func (c *ctx) Error() *errors.Error {
	if c.err == nil {
		c.err = errors.New()
	}
	return c.err
}

func (c *ctx) Validate(i interface{}) error {
	err := validate.Validate(i)
	if err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			for _, err := range ve {
				c.Error().AddError(err.Field(), err.Tag())
			}
		}
	}
	return err
}

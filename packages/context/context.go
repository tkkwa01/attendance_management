package context

import (
	"attendance-management/packages/errors"
	"attendance-management/packages/validate"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Context interface {
	Error() *errors.Error
	Validate(i interface{}) error
	DB() *gorm.DB
	RequestID() string
}

type ctx struct {
	err   *errors.Error
	getDB func() *gorm.DB
	db    *gorm.DB
	id    string
	uid   uint
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

func New(c *gin.Context, getDB func() *gorm.DB) Context {
	requestID := c.GetHeader("X-Request-Id")
	if requestID == "" {
		requestID = uuid.New().String()
	}

	var uid uint
	claimsInterface, ok := c.Get("claims")
	if ok {
		if uidInterface, ok := claimsInterface.(map[string]interface{})["uid"]; ok {
			uid = uint(uidInterface.(float64))
		}
	}

	var err *errors.Error
	if errInterface, ok := c.Get("error"); ok {
		err = errInterface.(*errors.Error)
	}

	return &ctx{
		id:    requestID,
		getDB: getDB,
		uid:   uid,
		err:   err,
	}
}

func (c *ctx) RequestID() string {
	return c.id
}

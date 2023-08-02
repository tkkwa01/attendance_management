package validate

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func Validate(i interface{}) error {
	return validate.Struct(i)
}

//interface{}を使うことでどんな型でも受け入れる

package middleware

import (
	"github.com/gookit/validate"
)

// var TrimTag validator.Func = func(fl validator.FieldLevel) bool {
// 	s, ok := fl.Field().Interface().(string)
// 	if !ok {
// 		return false

// 	}
// 	s = strings.TrimSpace(s)
// 	fl.Field().SetString(s)
// 	return true
// }

type CustomValidator struct { } 

func (c *CustomValidator) ValidateStruct(ptr any) error {
	v := validate.Struct(ptr)
	v.Validate()

	if v.Errors.Empty() {
		return nil
	}

	return v.Errors.ErrOrNil()
	// return v.Errors
}

func (c *CustomValidator) Engine() any {
	return nil
}
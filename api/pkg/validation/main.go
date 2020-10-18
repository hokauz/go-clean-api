package validation

import (
	"errors"
	"reflect"
	"strings"

	validator "github.com/go-playground/validator/v10"
)

var check *validator.Validate

// Test -
func Test(data interface{}) error {
	check = validator.New()

	check.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})

	err := check.Struct(data)
	if err != nil {
		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return nil
		}

		var arrErr []string
		for _, err := range err.(validator.ValidationErrors) {
			arrErr = append(arrErr, "Missing field "+err.Field()+" with type "+err.Kind().String())
		}

		str := strings.Join(arrErr, ", ")
		// from here you can create your own error messages in whatever language you wish
		return errors.New(str)
	}
	return nil
}

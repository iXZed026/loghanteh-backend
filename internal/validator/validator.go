package validator

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

func Register(v *validator.Validate) error {
	return v.RegisterValidation(
		"price5000",
		validatePrice5000,
	)
}

func validatePrice5000(fl validator.FieldLevel) bool {
	value := fl.Field()

	if value.Kind() != reflect.Int {
		return false
	}

	price := value.Int()

	return price%5000 == 0

}

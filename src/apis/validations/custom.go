package validations

import (
	"log"
	"regexp"

	"github.com/go-playground/validator/v10"
)

func IranianMobileValidator(fl validator.FieldLevel) bool {
	mobile, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}
	ok, err := regexp.MatchString(`^(0|0098|\+98)9(0[1-5]|[1 3]\d|2[0-2]|98)\d{7}$`, mobile)
	if err != nil {
		log.Print(err.Error())
		return false
	}
	if !ok {
		return false
	}
	return true

}

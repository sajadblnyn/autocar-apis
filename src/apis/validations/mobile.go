package validations

import (
	"github.com/go-playground/validator/v10"
	"github.com/sajadblnyn/autocar-apis/common"
)

func IranianMobileValidator(fl validator.FieldLevel) bool {
	mobile, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	return common.ValidateIranianMobile(mobile)

}

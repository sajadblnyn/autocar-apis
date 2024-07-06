package common

import (
	"regexp"

	"github.com/sajadblnyn/autocar-apis/config"
	"github.com/sajadblnyn/autocar-apis/pkg/logging"
)

const mobile_regex_pattern string = `^(0|0098|\+98)9(0[1-5]|[1 3]\d|2[0-2]|98)\d{7}$`

var logger logging.Logger = logging.NewLogger(config.GetConfig())

func ValidateIranianMobile(mobile string) bool {
	ok, err := regexp.MatchString(mobile_regex_pattern, mobile)
	if err != nil {
		logger.Error(logging.Validation, logging.MobileValidation, err.Error(), nil)
		return false
	}
	return ok
}

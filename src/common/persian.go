package common

import (
	"log"
	"regexp"
)

const mobile_regex_pattern string = `^(0|0098|\+98)9(0[1-5]|[1 3]\d|2[0-2]|98)\d{7}$`

func ValidateIranianMobile(mobile string) bool {
	ok, err := regexp.MatchString(mobile_regex_pattern, mobile)
	if err != nil {
		log.Print(err.Error())
		return false
	}
	return ok
}

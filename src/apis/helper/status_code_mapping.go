package helper

import (
	"net/http"

	"github.com/sajadblnyn/autocar-apis/pkg/service_errors"
)

var OtpStatusCodes map[string]int = map[string]int{
	service_errors.OtpUsed:     http.StatusConflict,
	service_errors.NotValidOtp: http.StatusConflict,
	service_errors.OtpExists:   http.StatusBadRequest,

	service_errors.EmailExists:    http.StatusConflict,
	service_errors.UsernameExists: http.StatusConflict,

	service_errors.RecordNotFound:   http.StatusNotFound,
	service_errors.PermissionDenied: http.StatusForbidden,
}

func TranslateErrorToStatusCode(err error) int {
	code, exists := OtpStatusCodes[err.Error()]
	if !exists {
		return http.StatusInternalServerError
	}
	return code
}

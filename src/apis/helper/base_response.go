package helper

import "github.com/sajadblnyn/autocar-apis/apis/validations"

type BaseHttpResponse struct {
	Result           any                            `json:"result"`
	Success          bool                           `json:"success"`
	ValidationErrors *[]validations.ValidationError `json:"validationErrors"`
	Error            any                            `json:"error"`
	ResultCode       int                            `json:"resultCode"`
}

func GenerateBaseResponse(result any, success bool, resultCode int) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result:     result,
		Success:    success,
		ResultCode: resultCode,
	}
}

func GenerateBaseResponseWithError(result any, success bool, resultCode int, err error) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result:     result,
		Success:    success,
		ResultCode: resultCode,
		Error:      err.Error(),
	}
}

func GenerateBaseResponseWithValidationErrors(result any, success bool, resultCode int, err error) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result:           result,
		Success:          success,
		ResultCode:       resultCode,
		Error:            err.Error(),
		ValidationErrors: validations.GetValidationErrors(err),
	}
}

package helpers

import "strings"

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
	Result  interface{} `json:"result"`
}

type EmptyResponse struct{}

func SuccessResponse(message string, result interface{}) Response {
	res := Response{
		Success: true,
		Message: message,
		Result:  result,
	}

	return res
}

func ErrorResponse(message string, err string, data interface{}) Response {
	splittedError := strings.Split(err, "\n")
	res := Response{
		Success: false,
		Message: message,
		Error:   splittedError,
	}

	return res
}

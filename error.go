package ups

import "fmt"

type ErrorResponse struct {
	Errors []Error `json:"errors"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e ErrorResponse) Error() string {
	errorString := ""

	for _, v := range e.Errors {
		errorString += fmt.Sprintf("%s:%s", v.Code, v.Message)
	}

	return errorString
}

package httperors

import (
	// "fmt"
	"net/http"
)

type HttpErr interface {
	Message() string
	Code() int
	Errors() string
	Noerror() bool
}

// func (e HttpError) Error() string {
//   return fmt.Sprintf("Message: %s - status: %d - Error: %s", e.Message, e.Code, e.Error)
// }

func (e HttpError) Message() string {
	return e.message
}
func (e HttpError) Code() int {
	return e.code
}
func (e HttpError) Errors() string {
	return e.errors
}
func (e HttpError) Noerror() bool {
	return e.results
}
func NewHttpError(message, errors string, code int) HttpErr {
	return HttpError{
		message: message,
		code:    code,
		errors:  errors,
	}
}
func NewBadRequestError(message string) HttpErr {
	return HttpError{
		message: message,
		code:    http.StatusBadRequest,
		errors:  "Bad Request",
	}
}
func NewAnuthorizedError(message string) HttpErr {
	return HttpError{
		message: message,
		code:    http.StatusUnauthorized,
		errors:  "Unauthorized",
	}
}
func NewNotFoundError(message string) HttpErr {
	return HttpError{
		message: message,
		code:    http.StatusNotFound,
		errors:  "Not Found",
	}
}
func ValidStringNotEmpty(str string) HttpErr {
	if str == "" {
		return HttpError{
			message: "The string must not be empty!",
			code:    http.StatusNotFound,
			errors:  "The string is Empty!!!",
			results: true,
		}
	}
	return HttpError{
		results: false,
	}
}

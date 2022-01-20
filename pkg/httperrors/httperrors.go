package httperrors

import (
	"errors"
	"net/http"
)

var (
	ErrBadRequest          = errors.New("Bad request")
	ErrNotFound            = errors.New("Not Found")
	ErrNotRequiredFields   = errors.New("No such required fields")
	ErrBadQueryParams      = errors.New("Invalid query params")
	ErrInternalServerError = errors.New("Internal Server Error")
)

type RestError interface {
	Status() int
	Error() string
	Causes() interface{}
}

type AppError struct {
	ErrStatus  int         `json:"status,omitempty"`
	ErrMessage string      `json:"message,omitempty"`
	ErrCauses  interface{} `json:"-"`
}

func (e AppError) Status() int {
	return e.ErrStatus
}

func (e AppError) Error() string {
	return e.ErrMessage
}

func (e AppError) Causes() interface{} {
	return e.ErrCauses
}

func NewRestError(status int, message string, causes interface{}) RestError {
	return AppError{
		ErrStatus:  status,
		ErrMessage: message,
		ErrCauses:  causes,
	}
}

func NewBadRequestError(causes interface{}) RestError {
	return AppError{
		ErrStatus:  http.StatusBadRequest,
		ErrMessage: ErrBadRequest.Error(),
		ErrCauses:  causes,
	}
}

func NewNotFoundError(causes interface{}) RestError {
	return AppError{
		ErrStatus:  http.StatusUnauthorized,
		ErrMessage: ErrNotFound.Error(),
		ErrCauses:  causes,
	}
}

func NewInternalServerError(causes interface{}) RestError {
	return AppError{
		ErrStatus:  http.StatusInternalServerError,
		ErrMessage: ErrInternalServerError.Error(),
		ErrCauses:  causes,
	}
}

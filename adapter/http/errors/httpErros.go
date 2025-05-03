package errors

import "net/http"

type GoMarryError struct {
	StatusCode int    `json:"-"`
	Message    string `json:"message"`
}

func (e *GoMarryError) Error() string {
	return e.Message
}

func NewBadRequest(message string) *GoMarryError {
	return &GoMarryError{
		StatusCode: http.StatusBadRequest,
		Message:    message,
	}
}

func NewInternal(message string) *GoMarryError {
	return &GoMarryError{
		StatusCode: http.StatusInternalServerError,
		Message:    message,
	}
}

func NewConflict(message string) *GoMarryError {
	return &GoMarryError{
		StatusCode: http.StatusConflict,
		Message:    message,
	}
}

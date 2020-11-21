package errors

import (
	"net/http"
)

type RestErrors struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewBadRequestError(message string) *RestErrors {
	return &RestErrors{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad request",
	}
}

func NewInternalServerError(message string) *RestErrors {
	return &RestErrors{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal server error",
	}
}

func NewNotFoundError(message string) *RestErrors {
	return &RestErrors{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not found",
	}
}

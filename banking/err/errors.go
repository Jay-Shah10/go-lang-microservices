package err

import "net/http"

type AppErrors struct {
	Code    int    `json: "omitempty"`
	Message string `json: "message"`
}

func (e AppErrors) AsMessage() *AppErrors {
	return &AppErrors{
		Message: e.Message,
	}
}

func NewNotFoundError(message string) *AppErrors {
	return &AppErrors{
		Message: message,
		Code:    http.StatusNotFound,
	}
}

func NewUnexpectedError(message string) *AppErrors {
	return &AppErrors{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}

package httperror

import "net/http"

const (
	UndefinedErrorCode = "ERR-C-0000"
	PersonNotFound     = "ERR-T-0001"
	PersonCreateError  = "ERR-T-0002"
)

var errors = map[string]HttpError{
	PersonNotFound: {
		StatusCode:  http.StatusInternalServerError,
		Description: "Person not found.",
	},
	PersonCreateError: {
		StatusCode:  http.StatusInternalServerError,
		Description: "Person create error.",
	},
}

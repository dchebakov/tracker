package response

import (
	"net/http"

	"github.com/dchebakov/tracker/pkg/httperrors"
)

type Response struct {
	Status  string  `json:"status"`
	Message *string `json:"message,omitempty"`
}

func Ok(message *string) (int, Response) {
	return http.StatusOK, Response{
		Status:  "ok",
		Message: message,
	}
}

func Error(err error) (int, Response) {
	if rest, ok := err.(httperrors.RestError); ok {
		status := rest.Status()
		message := rest.Error()

		return status, Response{
			Status:  "error",
			Message: &message,
		}
	}

	general := httperrors.NewInternalServerError(err)
	status := general.Status()
	message := general.Error()

	return status, Response{
		Status:  "error",
		Message: &message,
	}
}

package response

import (
	"net/http"

	"github.com/dchebakov/tracker/pkg/httperrors"
)

type Response struct {
	Status  string      `json:"status"`
	Message *string     `json:"message,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
}

func Ok(payload interface{}) (int, Response) {
	return http.StatusOK, Response{
		Status:  "ok",
		Payload: payload,
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

package response

import (
	"encoding/json"
	"net/http"

	"github.com/RupeshMahanta1994/flightops/shared/errors"
)

func Error(
	w http.ResponseWriter,
	err *errors.AppError,
) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.HTTPStatus)

	return json.NewEncoder(w).Encode(ErrorResponse{
		Success: false,
		Message: err.Message,
		Error: ErrorBody{
			Code: err.Code,
		},
	})
}

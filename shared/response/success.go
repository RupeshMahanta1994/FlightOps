package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Success()
func Success(w http.ResponseWriter, status int, message string, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	resp := SuccessResponse{
		Success: true,
		Message: message,
		Data:    data,
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		fmt.Println("Error in decoding response ", err)
		http.Error(
			w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
	}
}

// Created()

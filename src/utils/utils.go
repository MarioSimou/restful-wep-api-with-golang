package utils

import (
	"encoding/json"
	"models"
	"net/http"
)

type WrappedResponseSuccessOne struct {
	Status  int16       `json:"status"`
	Success bool        `json:"success"`
	Data    models.User `json:"data"`
}

type WrappedResponseSuccessMany struct {
	Status  int16         `json:"status"`
	Success bool          `json:"success"`
	Data    []models.User `json:"data"`
}

type WrappedResponseError struct {
	Status  int16  `json:"status"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func ErrorResponse(e error, w http.ResponseWriter) {
	// HTTP Response Headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	json.NewEncoder(w).Encode(WrappedResponseError{
		Status:  400,
		Success: false,
		Message: e.Error(),
	})
	return
}

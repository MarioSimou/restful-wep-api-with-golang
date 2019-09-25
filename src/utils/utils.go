package utils

import (
	"encoding/json"
	"models"
	"net/http"
)

type WrappedResponseSuccessOne struct {
	Status  int       `json:"status"`
	Success bool        `json:"success"`
	Data    models.User `json:"data"`
}

type WrappedResponseSuccessMany struct {
	Status  int         `json:"status"`
	Success bool          `json:"success"`
	Data    []models.User `json:"data"`
}

type WrappedResponseError struct {
	Status  int  `json:"status"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func ErrorResponse(status int, message string, w http.ResponseWriter) {
	// HTTP Response Headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(WrappedResponseError{
		Status:  status,
		Success: false,
		Message: message,
	})
}
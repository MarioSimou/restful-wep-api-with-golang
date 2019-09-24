package routes

import (
	"encoding/json"
	m "models"
	"net/http"
	"time"
)

// GET /api/users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	u := m.User{
		ID:         1,
		Username:   "john",
		Email:      "johndoe@gmail.com",
		Password:   "1234",
		Role:       "basic",
		CreatedAt:  time.Now().Format("UnixDate"),
		ModifiedAt: time.Now().Format("UnixDate"),
	}

	json.NewEncoder(w).Encode(m.WrappedResponseSuccessOne{
		Status:  200,
		Success: true,
		Data:    u,
	})

}

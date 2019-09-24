package routes

import (
	"encoding/json"
	"models"
	"net/http"
	"utils"

	"github.com/jinzhu/gorm"
)

func RenderRoute(route string, db *gorm.DB) func(http.ResponseWriter, *http.Request) {
	// GET /api/users
	getUsers := func(w http.ResponseWriter, r *http.Request) {
		var users []models.User

		// get users
		db.Find(&users)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(utils.WrappedResponseSuccessMany{
			Status:  200,
			Success: true,
			Data:    users,
		})

	}

	// POST /api/users
	postUser := func(w http.ResponseWriter, r *http.Request) {
		var user models.User

		if e := json.NewDecoder(r.Body).Decode(&user); e != nil {
			utils.ErrorResponse(e, w)
		}

		// stores the user in the database
		db.Create(&user)

		// HTTP Response Headers
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)

		// JSON response
		json.NewEncoder(w).Encode(utils.WrappedResponseSuccessOne{
			Status:  200,
			Success: true,
			Data:    user,
		})
	}

	switch route {
	case "getUsers":
		return getUsers
	case "postUser":
		return postUser
	default:
		return getUsers
	}
}

package routes

import (
	"encoding/json"
	"models"
	"net/http"
	"utils"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

const (
	GET_USERS = "getUsers"
	GET_USER = "getUser"
	POST_USER = "postUser"
	DELETE_USER = "deleteUser"
	PUT_USER = "putUser" 
)

func FetchRoute(route string, db *gorm.DB) func(http.ResponseWriter, *http.Request) {
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

	// GET /api/users/:id
	getUser := func(w http.ResponseWriter, r *http.Request){
		var user models.User
		// extract URI parameters and validates them
		params := mux.Vars(r)
		id := params["id"]
		if id == "" {
			utils.ErrorResponse(404,"Undefined user id", w)
			return
		}

		// HTTP 404 - Not Found
		if db.Find(&user, id); user.Id == 0 {
			utils.ErrorResponse(404,"User does not exist", w)
			return
		}

		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(utils.WrappedResponseSuccessOne{
			Status: 200,
			Success: false,
			Data: user,
		})
	}

	// POST /api/users
	postUser := func(w http.ResponseWriter, r *http.Request) {
		var user models.User

		if e := json.NewDecoder(r.Body).Decode(&user); e != nil {
			utils.ErrorResponse(400,e.Error(), w)
		}

		// stores the user in the database
		db.Create(&user)

		// HTTP Response Headers
		w.Header().Set("Location", r.URL.Path + "/" + strconv.Itoa(user.Id))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		json.NewEncoder(w).Encode(utils.WrappedResponseSuccessOne{
			Status:  201,
			Success: true,
			Data:    user,
		})
	}

	putUser := func(w http.ResponseWriter, r *http.Request){
		var user models.User
		var updUser models.User
		params := mux.Vars(r)
		id := params["id"]
		// HTTP 400 - Bad Request
		if id == "" {
			utils.ErrorResponse(400, "Unable to parse URI parameters", w )
			return
		}

		if e := json.NewDecoder(r.Body).Decode(&updUser); e != nil {
			utils.ErrorResponse(400, "Unable to parse http body", w )
			return
		}

		// HTTP 404 - Not Found
		if db.Where("id=?", id ).First(&user); user.Id == 0 {
			utils.ErrorResponse(404, "User does not exists", w )
			return
		}

		// procceeds updating the user
		db.Model(&user).Updates(updUser)


		// HTTP 200 - OK
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(utils.WrappedResponseSuccessOne{
			Status: 200,
			Success: true,
			Data: user,
		})
	}

	// DELETE /api/users/:id
	deleteUser := func(w http.ResponseWriter, r *http.Request){
		var user models.User
		params := mux.Vars(r)
		id := params["id"]
		if id == "" {
			utils.ErrorResponse(400,"Unable to parse URI parameters", w)
			return
		}

		// HTTP 404 - Not Found
		if db.Where("id=?", id).First(&user); user.Id == 0 {
			utils.ErrorResponse(404,"User does not exists",w)
			return
		}

		// if the user exists is deleted
		db.Delete(user)

		// HTTP 200 - OK
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(utils.WrappedResponseSuccessOne{
			Status: 200,
			Success: true,
			Data: user,
		})
	}

	switch route {
	case GET_USERS: return getUsers
	case GET_USER:  return getUser
	case POST_USER: return postUser
	case PUT_USER:  return putUser
	case DELETE_USER: return deleteUser
	default: return nil
	}
}

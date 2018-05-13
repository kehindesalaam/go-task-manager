package controllers

import (
	"encoding/json"
	"net/http"
	"taskmanager/common"
	"taskmanager/data"
	"taskmanager/models"
)

//Handler for HTTP Post - "/users/register
//Add a new User Document
func Register(w http.ResponseWriter, r *http.Request) {
	var dataResource UserResource
	//Decode the incoming User json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid User data", 500)
		return
	}
	user := &dataResource.Data
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("users")
	repo := &data.UserRepository{c}
	//Insert User document
	repo.CreateUser(user)
	//Clean-up the hash password to eliminate it from the response
	user.HashPassword = nil
	if j, err := json.Marshal(UserResource{Data: *user}); err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occured", 500)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(j)
	}
}

//Handler for HTTP Post - "/users/login
//Authenticate with username and password
func Login(w http.ResponseWriter, r *http.Request) {
	var dataResource LoginResource
	var token string
	//Decode the incoming Login json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid Login data", 401)
		return
	}
	loginModel := dataResource.Data
	loginUser := models.User{
		Email:    loginModel.Email,
		Password: loginModel.Password,
	}
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("users")
	repo := &data.UserRepository{C: c}

	//Authenticate the logged in user
	if user, err := repo.Login(loginUser); err != nil {
		common.DisplayAppError(w, err, "Invalid login credentials", 401)
		return
	} else {
		//Generate JWT token
		token, err = common.GenerateJWT(user.Email, user.Id, "member")
		if err != nil {
			common.DisplayAppError(w, err, "Error while generating the access token", 500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		user.HashPassword = nil
		authUser := AuthUserModel{
			User:  user,
			Token: token,
		}
		j, err := json.Marshal(AuthUserResource{Data: authUser})
		if err != nil {
			common.DisplayAppError(
				w,
				err,
				"An unexpected error has occurred",
				500,
			)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}

}

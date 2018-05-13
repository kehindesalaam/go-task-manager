package controllers

import "taskmanager/models"

type (
	//For Post - /user/register
	UserResource struct {
		Data models.User `json:"data"`
	}

	//For Post - /user/login
	LoginResource struct {
		Data LoginModel `json:"data"`
	}
	//Response for authorized user Post - /user/login
	AuthUserResource struct {
		Data AuthUserModel `json:"data"`
	}

	//Model for Authentication
	LoginModel struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	//Model for authorized user with access token
	AuthUserModel struct {
		User  models.User `json:"user"`
		Token string      `json:"token"`
	}
)

//tasks
type (
	//For Post/Put - /tasks
	//Foe Get - /tasks/id
	TaskResource struct {
		Data models.Task `json:"data"`
	}
	//For Get - /tasks
	TasksResource struct {
		Data []models.Task `json:"data"`
	}
)

//notes
type (
	NoteResource struct {
		Data models.TaskNote `json:"data"`
	}
	// For Get - /notes
	// For /notes/tasks/id
	NotesResource struct {
		Data []models.TaskNote `json:"data"`
	}
)

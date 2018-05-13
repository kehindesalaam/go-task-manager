package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type (
	User struct {
		Id           bson.ObjectId `bson:"_id,omitempty" json:"id"`
		FirstName    string        `json:"first_name"`
		LastName     string        `json:"last_name"`
		Email        string        `json:"email"`
		Password     string        `json:"password,omitempty"`
		HashPassword []byte        `json:"hash_password,omitempty"`
	}
	Task struct {
		Id          bson.ObjectId `bson:"_id,omitempty" json:"id"`
		CreatedBy   string        `json:"created_by"`
		Name        string        `json:"name"`
		Description string        `json:"description"`
		CreatedOn   time.Time     `json:"created_on,omitempty"`
		Due         time.Time     `json:"due,omitempty"`
		Status      string        `json:"status,omitempty"`
		Tags        []string      `json:"tags,omitempty"`
	}
	TaskNote struct {
		Id          bson.ObjectId `bson:"_id,omitempty" json:"id"`
		TaskId      bson.ObjectId `json:"task_id"`
		Description string        `json:"description"`
		CreatedOn   time.Time     `json:"created_on,omitempty"`
	}
)

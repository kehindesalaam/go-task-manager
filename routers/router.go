package routers

import (
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	//Routes for the user entity
	router = SetUserRoutes(router)
	//Routes for teh task entity
	router = SetTaskRoutes(router)
	//Routes for the TaskNote entity
	router = SetNoteRoutes(router)
	return router
}

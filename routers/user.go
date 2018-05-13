package routers

import (
	"github.com/gorilla/mux"
	"net/http"
	"taskmanager/controllers"
)

func SetUserRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/users/register", controllers.Register).Methods(http.MethodPost)
	router.HandleFunc("/users/login", controllers.Login).Methods(http.MethodPost)
	return router
}

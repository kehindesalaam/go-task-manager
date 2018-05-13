package routers

import (
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"taskmanager/common"
	"taskmanager/controllers"
	"github.com/gorilla/context"
)

func SetTaskRoutes(router *mux.Router) *mux.Router {
	taskRouter := mux.NewRouter()
	commonHandlers := alice.New(context.ClearHandler,common.Authorize)
	taskRouter.Handle("/tasks", commonHandlers.ThenFunc(controllers.CreateTask)).Methods("POST")
	taskRouter.Handle("/tasks/{id}", commonHandlers.ThenFunc(controllers.UpdateTask)).Methods("PUT")
	taskRouter.Handle("/tasks", commonHandlers.ThenFunc(controllers.GetTasks)).Methods("GET")
	taskRouter.Handle("/tasks/{id}", commonHandlers.ThenFunc(controllers.GetTaskById)).Methods("GET")
	taskRouter.Handle("/tasks/users/{id}", commonHandlers.ThenFunc(controllers.GetTasksByUser)).Methods("GET")
	taskRouter.Handle("/tasks/{id}", commonHandlers.ThenFunc(controllers.DeleteTask)).Methods("DELETE")

	router.PathPrefix("/tasks").Handler(taskRouter)
	return router
}

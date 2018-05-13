package routers

import (
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"taskmanager/common"
	"taskmanager/controllers"
	"github.com/gorilla/context"
)

func SetNoteRoutes(router *mux.Router) *mux.Router {
	noteRouter := mux.NewRouter()
	commonHandlers := alice.New(context.ClearHandler,common.Authorize)
	noteRouter.Handle("/notes", commonHandlers.ThenFunc(controllers.CreateNote)).Methods("POST")
	noteRouter.Handle("/notes/{id}", commonHandlers.ThenFunc(controllers.UpdateNote)).Methods("PUT")
	noteRouter.Handle("/notes/{id}", commonHandlers.ThenFunc(controllers.GetNoteById)).Methods("GET")
	noteRouter.Handle("/notes", commonHandlers.ThenFunc(controllers.GetNotes)).Methods("GET")
	noteRouter.Handle("/notes/tasks/{id}", commonHandlers.ThenFunc(controllers.GetNotesByTask)).Methods("GET")
	noteRouter.Handle("/notes/{id}", commonHandlers.ThenFunc(controllers.DeleteNote)).Methods("DELETE")

	router.PathPrefix("/notes").Handler(noteRouter)
	return router
}

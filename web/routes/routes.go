// Package routes ...
package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	conf "github.com/wkozyra95/go-web-starter/config"
	"net/http"
)

const (
	paramProjectId = "projectId"
	paramNoteId    = "noteId"
)

func SetupHandler(router *mux.Router, config conf.Config) error {
	context, contextErr := setupServerContext(config)
	if contextErr != nil {
		return contextErr
	}
	setupAuthHandler(router.PathPrefix("/auth").Subrouter(), context)
	setupProjectsHandler(router.PathPrefix("/projects").Subrouter(), context)
	setupNotesHandler(router.PathPrefix("/notes").Subrouter(), context)
	return nil
}

func setupAuthHandler(router *mux.Router, context serverContext) {
	router.Handle("/login", loginHandler(context)).Methods(http.MethodPost)
	router.Handle("/register", loginHandler(context)).Methods(http.MethodPost)
}

func setupProjectsHandler(router *mux.Router, context serverContext) {
	router.Handle(
		"",
		getProjectsHandler(context),
	).Methods(http.MethodGet)

	router.Handle(
		fmt.Sprintf("/{%s}", paramProjectId),
		getProjectHandler(context),
	).Methods(http.MethodGet)

	router.Handle(
		"",
		createProjectHandler(context),
	).Methods(http.MethodPost)

	router.Handle(
		fmt.Sprintf("/{%s}", paramProjectId),
		updateProjectHandler(context),
	).Methods(http.MethodPut)

	router.Handle(
		fmt.Sprintf("/{%s}", paramProjectId),
		deleteProjectHandler(context),
	).Methods(http.MethodDelete)
}

func setupNotesHandler(router *mux.Router, context serverContext) {
	router.Handle(
		fmt.Sprintf("/{%s}", paramProjectId),
		getNotesHandler(context),
	).Methods(http.MethodGet)

	router.Handle(
		fmt.Sprintf("/{%s}/{%s}", paramProjectId, paramNoteId),
		getNoteHandler(context),
	).Methods(http.MethodGet)

	router.Handle(
		fmt.Sprintf("/{%s}", paramProjectId),
		createNoteHandler(context),
	).Methods(http.MethodPost)

	router.Handle(
		fmt.Sprintf("/{%s}/{%s}", paramProjectId, paramNoteId),
		updateNoteHandler(context),
	).Methods(http.MethodPut)

	router.Handle(
		fmt.Sprintf("/{%s}/{%s}", paramProjectId, paramNoteId),
		deleteNoteHandler(context),
	).Methods(http.MethodDelete)
}

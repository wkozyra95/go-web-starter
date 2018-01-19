package routes

import (
	"net/http"
)

type getProjectsHandler serverContext

func (ph getProjectsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusTeapot)
}

type getProjectHandler serverContext

func (ph getProjectHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusTeapot)
}

type createProjectHandler serverContext

func (ph createProjectHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusTeapot)
}

type updateProjectHandler serverContext

func (ph updateProjectHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusTeapot)
}

type deleteProjectHandler serverContext

func (ph deleteProjectHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusTeapot)
}

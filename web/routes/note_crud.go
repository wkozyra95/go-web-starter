package routes

import (
	"net/http"
)

type getNotesHandler serverContext

func (nh getNotesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusTeapot)
}

type getNoteHandler serverContext

func (nh getNoteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusTeapot)
}

type createNoteHandler serverContext

func (nh createNoteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusTeapot)
}

type updateNoteHandler serverContext

func (nh updateNoteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusTeapot)
}

type deleteNoteHandler serverContext

func (nh deleteNoteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusTeapot)
}

package routes

import (
	"net/http"
)

type registerHandler serverContext

func (rh registerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusTeapot)
}

package routes

import (
	"net/http"
)

type loginHandler serverContext

func (lh loginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusTeapot)
}

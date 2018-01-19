// Package web ...
package web

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	conf "github.com/wkozyra95/go-web-starter/config"
	"github.com/wkozyra95/go-web-starter/web/routes"
)

func NewRouter(config conf.Config) (http.Handler, error) {
	router := mux.NewRouter()

	setupErr := routes.SetupHandler(config)
	if setupErr != nil {
		return nil, setupErr
	}

	return handlers.CORS(
		handlers.AllowedHeaders([]string{"content-type", "x-auth-token"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
	)(router), nil
}

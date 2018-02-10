// Package web ...
package web

import (
	"net/http"

	"github.com/go-chi/chi"
	conf "github.com/wkozyra95/go-web-starter/config"
	"github.com/wkozyra95/go-web-starter/model/mongo"
)

type handler struct {
	jwt           *jwtProvider
	config        *conf.Config
	dbCreatorFunc func() mongo.DB
}

func setupRoutes(h *handler, db dbProvider) (http.Handler, error) {
	router := chi.NewRouter()

	router.Use(db.middleware)
	router.Route("/auth", func(router chi.Router) {
		router.Post("/login", h.loginHandler)
		router.Post("/register", h.registerHandler)
	})
	router.Route("/projects", func(router chi.Router) {
		router.Use(h.jwt.middleware)

		router.Get("/", h.getProjectsHandler)
		router.Get("/{projectId}", h.getProjectHandler)
		router.Post("/", h.createProjectHandler)
		router.Put("/{projectId}", h.updateProjectHandler)
		router.Delete("/{projectId}", h.deleteProjectHandler)
	})
	return router, nil
}

package web

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	conf "github.com/wkozyra95/go-web-starter/config"
)

const (
	paramProjectId = "projectId"
	paramNoteId    = "noteId"
)

type requestCtx struct {
	server *serverContext
	chi    *chi.Context
}

type wrappedHandlerFunc = func(w http.ResponseWriter, r *http.Request, ctx requestCtx)

func setupRoutes(context serverContext, config conf.Config) (http.Handler, error) {
	router := chi.NewRouter()
	f := func(handler wrappedHandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			handler(w, r, requestCtx{
				server: &context,
				chi:    chi.RouteContext(r.Context()),
			})
		}
	}

	router.Route("/auth", func(router chi.Router) {
		router.Post("/login", f(loginHandler))
		router.Post("/register", f(loginHandler))
	})
	router.Route("/projects", func(router chi.Router) {
		routeRoot := "/"
		routeId := fmt.Sprintf("/{%s}", paramProjectId)
		router.Use(context.jwt.middleware)

		router.Get(routeRoot, f(getProjectsHandler))
		router.Get(routeId, f(getProjectHandler))
		router.Post(routeRoot, f(createProjectHandler))
		router.Put(routeId, f(updateProjectHandler))
		router.Delete(routeId, f(deleteProjectHandler))
	})
	return router, nil
}

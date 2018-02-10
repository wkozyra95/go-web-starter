// Package web ...
package web

import (
	"net/http"

	conf "github.com/wkozyra95/go-web-starter/config"
	"github.com/wkozyra95/go-web-starter/model/mongo"
)

var log = conf.NamedLogger("web")

// NewRouter ...
func NewRouter(config *conf.Config) (http.Handler, error) {
	dbCreatorFunc, dbErr := mongo.SetupDB(config)
	if dbErr != nil {
		log.Error(dbErr.Error())
		return nil, dbErr
	}

	jwt, jwtErr := newJwtProvider(config)
	if jwtErr != nil {
		log.Error(jwtErr.Error())
		return nil, dbErr
	}

	context := &handler{
		config:        config,
		jwt:           &jwt,
		dbCreatorFunc: dbCreatorFunc,
	}

	router, setupRoutesErr := setupRoutes(context, dbProvider(dbCreatorFunc))
	if setupRoutesErr != nil {
		log.Error(dbErr.Error())
		return nil, setupRoutesErr
	}

	return router, nil
}

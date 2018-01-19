package routes

import (
	conf "github.com/wkozyra95/go-web-starter/config"
)

type serverContext struct {
	config conf.Config
}

func setupServerContext(config conf.Config) (serverContext, error) {
	return serverContext{
		config: config,
	}, nil
}

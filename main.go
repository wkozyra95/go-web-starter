package main

import (
	"net/http"

	conf "github.com/wkozyra95/go-web-starter/config"
	"github.com/wkozyra95/go-web-starter/web"
)

var log = conf.NamedLogger("main")

func main() {
	config := conf.SetupConfig()
	handler, handlerErr := web.NewRouter(config)
	if handlerErr != nil {
		log.Error(handlerErr.Error())
		return
	}
	log.Info("Serving content on port 8080")
	http.ListenAndServe(":3001", handler)
}

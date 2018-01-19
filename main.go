package main

import (
	"net/http"

	conf "github.com/wkozyra95/go-web-starter/config"
	"github.com/wkozyra95/go-web-starter/web"
)

func main() {
	config := conf.SetupConfig()

	handler, handlerErr := web.NewRouter(config)
	if handlerErr != nil {
		return
	}
	http.ListenAndServe(":8080")
}

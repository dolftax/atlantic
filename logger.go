package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func middleware(router http.Handler, serverConfig ServerConfig) http.Handler {
	// Logger
	if serverConfig.Logger == true {
		return handlers.CombinedLoggingHandler(os.Stdout, router)
	}
	return router
}

package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func middleware(router http.Handler, serverConfig serverConfig) http.Handler {
	// Logger
	if serverConfig.logger == true {
		return handlers.CombinedLoggingHandler(os.Stdout, router)
	}
	return router
}

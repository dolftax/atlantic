package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/spf13/viper"
)

func middleware(router http.Handler) http.Handler {
	// Logger
	if (viper.GetBool("logger")) == true {
		return handlers.CombinedLoggingHandler(os.Stdout, router)
	}
	return router
}

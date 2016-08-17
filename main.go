package main

import (
	"log"
	"net/http"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/julienschmidt/httprouter"
)

type Config struct {
	Server struct {
		port   int
		logger bool
	}

	Broker struct {
		broker            string
		result_backend    string
		result_expires_in int
		exchange          string
		exchange_type     string
		default_queue     string
		binding_key       string
	}
}

func ReadConfig() Config {
	var configfile = "config.toml"
	_, err := os.Stat(configfile)
	if err != nil {
		log.Fatal("Config file is missing: ", configfile)
	}

	var config Config
	if _, err := toml.DecodeFile(configfile, &config); err != nil {
		log.Fatal(err)
	}

	return config
}

func main() {
	var config = ReadConfig()
	router := httprouter.New()
	router.GET("/*path", queuePass)

	log.Println("Atlantic server listening at port"+":", config.server.port)

	log.Fatal(http.ListenAndServe(":"+config.server.port, middleware(router)))
}

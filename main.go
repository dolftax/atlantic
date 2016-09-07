package main

import (
	"log"
	"net/http"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/julienschmidt/httprouter"
)

type ServerConfig struct {
	port   string `toml:"port"`
	logger bool   `toml:"logger"`
}

func init() {
	_, err := os.Stat("config.toml")
	if err != nil {
		log.Fatal("Config file is missing: ")
	}
}

func main() {
	// Server config parsing
	var serverConfig ServerConfig
	_, err := toml.DecodeFile("config.toml", &serverConfig)
	if err != nil {
		log.Fatal(err)
	}

	// Router
	router := httprouter.New()
	// TODO: Point to routing function when it is up // router.GET("/*path", )

	log.Println("Atlantic server listening at port", serverConfig.port)

	log.Fatal(http.ListenAndServe(":"+serverConfig.port, middleware(router, serverConfig)))
}

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

type BrokerConfig struct {
	broker        string `toml:"broker"`
	protocol      string `toml:"protocol"`
	expiry        int    `toml:"result_expires_in"`
	exchange_name string `toml:"exchange"`
	exchange_type string `toml:"exchange_type"`
	default_queue string `toml:"default_queue"`
	binding_key   string `toml:"binding_key"`
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
	router.GET("/*path", queuePass)

	log.Println("Atlantic server listening at port", serverConfig.port)

	log.Fatal(http.ListenAndServe(":"+serverConfig.port, middleware(router, serverConfig)))
}

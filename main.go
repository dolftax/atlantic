package main

import (
	"log"
	"net/http"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/gorilla/websocket"
)

type ServerConfig struct {
	port   string `toml:"port"`
	logger bool   `toml:"logger"`
}

func init() {
	_, err := os.Stat("config.toml")
	if err != nil {
		log.Fatal("Config file is missing")
	}
}

func main() {
	// Server config parsing
	var serverConfig ServerConfig
	_, err := toml.DecodeFile("config.toml", &serverConfig)
	if err != nil {
		log.Fatal(err)
	}

	http.Handle('/', func(w http.ResponseWriter, r *http.Request) {

		var upgrader = websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
		}
		// TODO: Pass on the connection to handler
	})
	
	log.Println("Atlantic server listening at port", serverConfig.port)
	log.Fatal(http.ListenAndServe(":"+serverConfig.port, middleware(router, serverConfig)))
}

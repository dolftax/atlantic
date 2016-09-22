package main

import (
	"log"
	"net"
	"net/http"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
)

type ServerConfig struct {
	Port   string `toml:"port"`
	Logger bool   `toml:"logger"`
}

func init() {
	_, err := os.Stat("config.toml")
	if err != nil {
		log.Fatal("Config file is missing")
	}
}

func upgrade_conn(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	conn_ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	} else {
		// Establish connection with docker daemon
		conn_docker, err := net.Dial("unix", "/var/run/docker.sock")
		if err != nil {
			log.Fatal(err)
		}

		// Pass connection object to message handler
		frame_handler(conn_ws, &conn_docker)
	}
}

func main() {
	// Server config parsing
	var serverConfig ServerConfig
	_, err := toml.DecodeFile("config.toml", &serverConfig)
	if err != nil {
		log.Fatal(err)
	}

	router := httprouter.New()

	// Establish websocket connection and serve the webapp
	router.GET("/", upgrade_conn)

	log.Println("Atlantic server listening at port", serverConfig.Port)
	log.Fatal(http.ListenAndServe(":"+serverConfig.Port, middleware(router, serverConfig)))
}

package main

import (
	"encoding/json"
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

// Docker connection object
var conn_docker net.Conn

func init() {
	// Check if config file exists
	_, err := os.Stat("config.toml")
	if err != nil {
		log.Fatal("Config file is missing", err)
	}
}

func upgrade_conn(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	// Define buffer size
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	conn_ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading connection", err)
	} else {
		// Establish connection with docker daemon
		conn_docker, err := net.Dial("unix", "/var/run/docker.sock")
		log.Println(conn_docker)
		if err != nil {
			// TODO: Make this optional. If unable to establish connection, request of engine type `docker` should be errored.
			log.Fatal("Error establishing connection with docker", err)
		}

		// Listen for websocket message frames
		for {
			message_type, request_frame, err := conn_ws.ReadMessage() // Frame type ignored
			if err != nil {
				log.Println("Error reading message frame", err)
			}

			if len(request_frame) == 0 {
				error_handler(conn_ws, messa "connection upgrade", 1000, "")
			} else {
				var request_obj map[string]string
				err := json.Unmarshal(request_frame, &request_obj)
				if err != nil {
					error_handler(conn_ws, "error parsing request JSON", 1001, "")
				}

				// Frame handler will take websocket connection and request object and takes ownership from here
				go frame_handler(conn_ws, request_obj)
			}
		}
	}
}

func main() {
	// Parse server configuration from `config.toml` file
	var serverConfig ServerConfig
	_, err := toml.DecodeFile("config.toml", &serverConfig)
	if err != nil {
		log.Fatal(err)
	}

	router := httprouter.New()

	// Establish websocket connection and serve the web application
	router.GET("/", upgrade_conn)

	log.Println("Atlantic server listening at port", serverConfig.Port)
	log.Fatal(http.ListenAndServe(":"+serverConfig.Port, middleware(router, serverConfig)))
}

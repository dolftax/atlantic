package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

type ResponseObj struct {
	Operation string
	Err       int
	Timestamp time.Time
	Result    string
}

func response_dispatcher(conn_ws *websocket.Conn, response_obj ResponseObj) {
	response, err := json.Marshal(response_obj)
	if err != nil {
		log.Println("Error mashaling response_obj JSON")
		return
	}

	err = conn_ws.WriteMessage(response)
	if err != nil {
		log.Println("Error sending response JSON as frame")
		return
	}
}

func response_handler(conn_ws *websocket.Conn, operation string, result string) {
	response_obj := ResponseObj{operation, 0, time.Now(), result}
	response_dispatcher(conn_ws, response_obj)
}

func error_handler(conn_ws *websocket.Conn, operation string, err int, result string) {

	response_obj := ResponseObj{operation, err, time.Now(), result}
	response_dispatcher(conn_ws, response_obj)
}

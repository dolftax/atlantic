package main

import (
	"github.com/gorilla/websocket"
	"log"
)

// type RequestObj struct {
// 	Engine    string
// 	Command   string
// 	Arguments []string
// }

func frame_handler(ws *websocket.Conn, request_obj map[string]string) {
	log.Println(request_obj)
}

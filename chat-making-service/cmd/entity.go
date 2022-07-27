package main

import (
	"github.com/gorilla/websocket"
)

type WebSocketConnection struct {
	*websocket.Conn
	username string
}

type SocketResponse struct {
	TimeHandle string
	RoomID     string
}

type Config struct {
	log Log
}

type Log struct{}

const (
	webPort = "8080"
)

var (
	queue    = make(chan *WebSocketConnection)
	upgrader websocket.Upgrader
)

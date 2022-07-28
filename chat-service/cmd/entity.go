package main

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Config struct {
	log Log
}

type WebSocketConnection struct {
	*websocket.Conn
	username string
}

type SocketMessage struct {
	Message string
}

type SocketResponse struct {
	From    string
	Message string
}

type Log struct{}

const (
	webPort = "8080"
)

var (
	upgrader  websocket.Upgrader
	chatRooms = make(map[string][]*WebSocketConnection)
	mu        sync.Mutex
)

package main

import (
	"log"

	"github.com/google/uuid"
)

func (app *Config) roomCreatorService(queue chan *WebSocketConnection) {
	conn1 := <-queue
	conn2 := <-queue

	roomID := uuid.New().String()

	if err := conn1.WriteJSON(SocketResponse{TimeHandle: "", RoomID: roomID}); err != nil {
		app.log.error(err.Error())
		return
	}

	if err := conn1.Close(); err != nil {
		app.log.error(err.Error())
		return
	}

	if err := conn2.WriteJSON(SocketResponse{TimeHandle: "", RoomID: roomID}); err != nil {
		app.log.error(err.Error())
		return
	}

	if err := conn2.Close(); err != nil {
		app.log.error(err.Error())
		return
	}

	app.roomCreatorService(queue)
}

func (app *Log) info(info string) {
	log.Printf("[INFO] %s", info)
}

func (app *Log) error(err string) {
	log.Printf("[ERROR] %s", err)
}

func (app *Log) panic(panic string) {
	log.Printf("[PANIC] %s", panic)
}

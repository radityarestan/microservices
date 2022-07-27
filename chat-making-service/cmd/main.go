package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

func init() {
	upgrader = websocket.Upgrader{}
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
}

func main() {
	app := Config{}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	app.log.info(fmt.Sprintf("Chat Making Service started at port: %s", webPort))

	go app.roomCreatorService(queue)

	if err := server.ListenAndServe(); err != nil {
		app.log.panic(err.Error())
	}
}

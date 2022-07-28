package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

func init() {
	upgrader = websocket.Upgrader{}
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
}

func main() {
	app := Config{}

	port := os.Getenv("PORT")
	if port == "" {
		port = fmt.Sprintf(":%s", webPort)
	}

	server := &http.Server{
		Addr:    port,
		Handler: app.routes(),
	}

	app.log.info(fmt.Sprintf("Chat Service started at port: %s", port))

	if err := server.ListenAndServe(); err != nil {
		app.log.panic(err.Error())
	}

}

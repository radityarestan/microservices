package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func (app *Config) findSomeone(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		app.log.error(err.Error())
		return
	}

	startTime := time.Now()

	username := r.URL.Query().Get("username")
	currentConn := &WebSocketConnection{conn, username}
	queue <- currentConn

	app.log.info(fmt.Sprintf("%s's connection has been added to queue", username))

	for {
		duration := time.Since(startTime)
		err := currentConn.WriteJSON(SocketResponse{
			TimeHandle: fmt.Sprintf("%.f", duration.Seconds()),
			RoomID:     "",
		})

		if err != nil {
			if strings.Contains(err.Error(), "use of closed network connection") {
				app.log.info(fmt.Sprintf("%s's connection has been removed from queue", username))
				return
			}
			app.log.error(err.Error())
		}

		time.Sleep(time.Second)
	}

}

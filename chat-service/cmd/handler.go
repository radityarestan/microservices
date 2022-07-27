package main

import (
	"net/http"
	"strings"
)

func (app *Config) chat(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		app.log.error(err.Error())
		return
	}

	roomID := r.URL.Query().Get("roomID")
	username := r.URL.Query().Get("username")
	currentConn := &WebSocketConnection{conn, username}

	mu.Lock()
	if _, ok := chatRooms[roomID]; !ok {
		chatRooms[roomID] = make([]*WebSocketConnection, 0)
	}
	chatRooms[roomID] = append(chatRooms[roomID], currentConn)
	mu.Unlock()

	for {
		payload := SocketMessage{}
		err := currentConn.ReadJSON(&payload)
		if err != nil {
			if strings.Contains(err.Error(), "websocket: close") {
				app.log.info(err.Error())
				return
			}

			app.log.error(err.Error())
			continue
		}

		broadcastMessage(currentConn, payload.Message, roomID)
	}
}

func broadcastMessage(currentConn *WebSocketConnection, message, roomID string) {
	for _, conn := range chatRooms[roomID] {
		if conn != currentConn {
			conn.WriteJSON(SocketResponse{Message: message, From: currentConn.username})
		}
	}
}

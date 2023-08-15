package controllers

import (
	"fmt"

	"golang.org/x/net/websocket"
)

type WSHandler struct {
}

type WSHandlerWrapper struct {
	wsHandler *WSHandler
}

func (w *WSHandlerWrapper) Handle(conn *websocket.Conn) {
	w.wsHandler.Handle(conn)
}

func NewWSHanlder() *WSHandlerWrapper {
	return &WSHandlerWrapper{}
}

func (w *WSHandler) Handle(conn *websocket.Conn) {
	defer conn.Close()

	uID := conn.Config().Location.Query().Get("userID")
	eType := conn.Config().Location.Query().Get("event_type")

	fmt.Printf("Websocket subscription active for user: %s and event type: %s", uID, eType)

	// some subscription service should take these params and add to internal storage for further actions

	for {
		// here is a logic to send events to client
	}

}

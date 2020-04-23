package main

import "github.com/gorilla/websocket"

type client struct {
	socker *websocket.Conn
	send   chan []byte
	room   *room
}

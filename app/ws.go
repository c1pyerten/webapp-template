package app

import "github.com/gorilla/websocket"

var Ws *websocket.Upgrader
const (
	bufferSize = 2048
)

func init() {
	Ws = &websocket.Upgrader{
		WriteBufferSize: bufferSize,
		ReadBufferSize: bufferSize,
		EnableCompression: true,
	}
}

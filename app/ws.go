package app

import "github.com/gorilla/websocket"

const (
	bufferSize = 2048
)

func newWs() *websocket.Upgrader {
	return &websocket.Upgrader{
		WriteBufferSize: bufferSize,
		ReadBufferSize: bufferSize,
		EnableCompression: true,
	}
}

// var Ws *websocket.Upgrader
// func init() {
// 	Ws = &websocket.Upgrader{
// 		WriteBufferSize: bufferSize,
// 		ReadBufferSize: bufferSize,
// 		EnableCompression: true,
// 	}
// }

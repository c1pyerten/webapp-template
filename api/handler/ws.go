package handler

import (
	"c1pherten/yet-webapp2/api"
	"c1pherten/yet-webapp2/app/ws"
	"c1pherten/yet-webapp2/container"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	bufferSize = 2048
)

type WsHandler struct {
	c         container.Container
	processor *ws.Processor
	upgrader  *websocket.Upgrader
}

func (w *WsHandler) HandleWs(ctx *gin.Context) {
	w.c.Logger().Info("ws connecting")
	conn, err := w.upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		w.c.Logger().Info(err)
		api.InternalError(ctx, err)
		return
	}
	defer conn.Close()

	for {
		typ, b, err := conn.ReadMessage()
		if err != nil {
			// w.c.Logger().Error(err, string(b))
			return
		}
		if typ != websocket.TextMessage {
			w.c.Logger().Error(typ, string(b))
			return
		}
		
		resp, err := w.processor.Handle(b)
		if err != nil {
			w.c.Logger().Error(err)
			// close connection
			return 
		}

		w.c.Logger().Info("resp", resp)
		if err := conn.WriteJSON(resp); err != nil {
			w.c.Logger().Error(err)
		}
		
	}

}

func NewWsHandler(c container.Container) *WsHandler {
	return &WsHandler{
		c:         c,
		processor: ws.NewProcessor(c),
		upgrader: &websocket.Upgrader{
			WriteBufferSize:   bufferSize,
			ReadBufferSize:    bufferSize,
			// EnableCompression: true,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}

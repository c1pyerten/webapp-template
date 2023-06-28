package handler

import (
	"c1pherten/yet-webapp2/api"
	"c1pherten/yet-webapp2/container"
	"c1pherten/yet-webapp2/service"

	"github.com/gin-gonic/gin"
)

type MessageHandler struct {
	c container.Container
	s *service.MsgService
}

func (h *MessageHandler) GetMsgById(ctx *gin.Context) {
	ctx.JSON(200, api.Response{
		Code: 0,
		Msg:  "todo",
		Data: nil,
	})

}

func NewMessageHandler(c container.Container) *MessageHandler {
	return &MessageHandler{
		c: c,
	}

}
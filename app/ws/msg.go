package ws

import (
	"c1pherten/yet-webapp2/container"
	"c1pherten/yet-webapp2/repository"
	"encoding/json"
)

type Msg struct {
	Type string `json:"type"`
	Data any    `json:"data"`
}

type MsgRaw struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

type PublicMsg struct {
	SourceID string `json:"sourceId"`
	Content  string `json:"content"`
	Ts       int64  `json:"ts"`
}

type PrivateMsg struct {
	SourceID string `json:"sourceID"`
	TargetID string `json:"targetID"`
	Content  string `json:"content"`
	Ts       int64  `json:"ts"`
}

type Tick struct{}

type msgHandler struct {
	c       container.Container
	msgRepo repository.MessageRepository
}

// func (h *msgHandler) handlePrivateMsg(a any) {
// 	m := a.(*PrivateMsg)
// 	h.c.Logger().Info(m)
// }

func (h *msgHandler) handlePrivateMsg() (*PrivateMsg, HandleFunc) {
	return nil, func(a any) (any, error) {
		m := a.(*PrivateMsg)
		h.c.Logger().Info("testing msg handler", m)
		return m, nil
	}
}

func (h *msgHandler) handlePublicMsg() (*PublicMsg, HandleFunc) {
	return nil, func(a any) (any, error) {
		m := a.(*PublicMsg)
		h.c.Logger().Info("testing msg handler pubmsg:", m)
		return m, nil
	}
}

func NewMsgHandler(c container.Container) *msgHandler {
	msgRepo := repository.NewMessageRepository(c)
	return &msgHandler{
		c:       c,
		msgRepo: msgRepo,

	}
}

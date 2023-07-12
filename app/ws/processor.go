package ws

import (
	"c1pherten/yet-webapp2/appctx"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

var (
	ErrUnregisteredMsg = errors.New("unregistered msg")
)

type Processor struct {
	c appctx.Container
	msgHandle msgHandler
	msgInfos map[string]*msgInfo
}

type msgInfo struct {
	msgType    reflect.Type
	msgHandler HandleFunc
}

type HandleFunc func(any) (any, error)


func (p *Processor) Register(m any) {
	if m == nil {
		panic("message should not be nil")
	}
	msgType := reflect.TypeOf(m)
	if msgType == nil || msgType.Kind() != reflect.Ptr {
		panic("message should be pointer")
	}
	msgId := msgType.Elem().Name()
	if msgId == "" {
		panic("unnamed message")
	}
	if _, ok := p.msgInfos[msgId]; ok {
		panic(fmt.Sprintf("duplicate msg %s", msgId))
	}

	p.msgInfos[msgId] = &msgInfo{
		msgType: msgType,
		msgHandler: func(any) (any, error) {
			return nil, nil
		},
	}
}

func (p *Processor) SetHandler(m any, h HandleFunc) {
	msgType := reflect.TypeOf(m)
	if msgType == nil || msgType.Kind() != reflect.Ptr {
		panic("pointer msg required")
	}
	msgId := msgType.Elem().Name()
	info, ok := p.msgInfos[msgId]
	if !ok {
		p.msgInfos[msgId] = &msgInfo{
			msgType:    msgType,
			msgHandler: h,
		}
		// panic(fmt.Sprintf("message %s not registered", msgId))
	} else {
		info.msgHandler = h
	}
}

func (p *Processor) Marshal(m any) ([]byte, error) {
	msgType := reflect.TypeOf(m)
	if msgType == nil {
		return nil, ErrUnregisteredMsg
	}
	msgID := msgType.Elem().Name()
	info, ok := p.msgInfos[msgID]
	_ = info
	if !ok {
		return nil, ErrUnregisteredMsg
	}
	b, err := json.Marshal(Msg{
		Type: msgID,
		Data: m,
	})
	if err != nil {
		return nil, err
	}
	return b, nil
}

// func (p *Processor) Unmarshal(data []byte) (any, error) {
func (p *Processor) Unmarshal(input []byte) (any, error) {
	var msgRaw MsgRaw
	if err := json.Unmarshal(input, &msgRaw); err != nil {
		return nil, err
	}
	info, ok := p.msgInfos[msgRaw.Type]
	if !ok {
		return nil, ErrUnregisteredMsg
	}

	data := reflect.New(info.msgType.Elem()).Interface()
	msg := Msg{
		Type: msgRaw.Type,
		Data: data,
	}
	return data, json.Unmarshal(input, &msg)
}

func getMsgID(m any) (string, error) {
	msgType := reflect.TypeOf(m)
	if msgType == nil || msgType.Kind() != reflect.Ptr {
		return "", fmt.Errorf("msg should be a non-nil pointer: %v", m)
	}
	msgID := msgType.Elem().Name()
	return msgID, nil
}

func (p *Processor) getMsgInfo(m any) (*msgInfo, error) {
	msgID, err := getMsgID(m)
	if err != nil {
		return nil, nil
	}
	info, ok := p.msgInfos[msgID]
	if !ok {
		return nil, ErrUnregisteredMsg
	}
	return info, nil
}

func (p *Processor) call(m any) (any, error) {
	info, err := p.getMsgInfo(m)
	if err != nil {
		return nil, err
	}

	if info.msgHandler != nil {
		return info.msgHandler(m)
	}
	return nil, ErrUnregisteredMsg
}

func (p *Processor) Handle(raw []byte) (any, error) {
	v, err := p.Unmarshal(raw)
	if err != nil {
		return nil, err
	}

	return p.call(v)
}

func NewProcessor(c appctx.Container) *Processor {
	p := &Processor{
		c:        c,
		msgInfos: make(map[string]*msgInfo),
	}
	mh := NewMsgHandler(c)

	
	// p.SetHandler((*PublicMsg)(nil), mh.handlePrivateMsg)
	p.SetHandler(mh.handlePrivateMsg())
	p.SetHandler(mh.handlePublicMsg())

	// p.SetHandler((*PublicMsg)(nil), func(a any) {
	// 	m := a.(*PublicMsg)
	// 	c.Logger().Info("in handler", m)
	// })


	return p
}
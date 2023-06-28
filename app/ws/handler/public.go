package handler

import (
	"c1pherten/yet-webapp2/repository"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PublicMsg struct {
	ID       primitive.ObjectID
	PublicID primitive.ObjectID
	Content  string
	Ts       int64
}

type mHandle struct {
	r repository.MessageRepository
}

// func (h *mHandle) FindMsg(id primitive.ObjectID) (*PublicMsg, error) {
// 	m, err := h.r.FindMessageByID(context.Background(), id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return m, nil
// }
func (h *mHandle) sendHandler() ()

func newHandler() *mHandle {
	return &mHandle{}
}

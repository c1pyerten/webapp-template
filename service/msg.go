package service

import (
	"c1pherten/yet-webapp2/appctx"
	"c1pherten/yet-webapp2/repository"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type MsgService struct{
	c appctx.Container
	msgRepo repository.MessageRepository
}

func (s *MsgService) FindMsgByID(id string) (*repository.Message, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	m, err := s.msgRepo.FindMessageByID(ctx, oid)
	if err != nil {
		return nil, err
	}
	
	return m, nil
}

func (s *MsgService) FindMsgByUserID(id int) ([]*repository.Message, error) {
	return nil, nil
}

func NewMsgService(c appctx.Container) *MsgService {
	return &MsgService{
		c: c,
	}
}
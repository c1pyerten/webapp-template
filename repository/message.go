package repository

import (
	"c1pherten/yet-webapp2/appctx"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const collName = "msgs"

type Message struct {
	ID       primitive.ObjectID `bson:"_id"`
	SourceID primitive.ObjectID `bson:"sourceID"`
	TargetID primitive.ObjectID `bson:"targetID"`
	PublicID primitive.ObjectID `bson:"publicID"`
	CreateAt time.Time          `bson:"createAt"`
	Ts       int64              `bson:"ts"`
	Content  string             `bson:"content"`
	Deleted  bool               `bson:"deleted"`
}

type MessageRepository interface {
	CreateMessage(context.Context, *Message) (primitive.ObjectID, error)
	FindMessageByID(context.Context, primitive.ObjectID) (*Message, error)
	FindMessagesByPrivateID(context.Context, primitive.ObjectID) ([]*Message, error)
	FindMessagesByPublicID(context.Context, primitive.ObjectID) ([]*Message, error)
	DeleteMsg(context.Context, primitive.ObjectID) error
}

type messageRepository struct {
	c appctx.Container
	repository
}

func (r *messageRepository) Coll() *mongo.Collection {
	return r.db.Collection(collName)
}

func (r *messageRepository) CreateMessage(ctx context.Context, m *Message) (primitive.ObjectID, error) {
	result, err := r.Coll().InsertOne(ctx, m)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return result.InsertedID.(primitive.ObjectID), nil
}

func (r *messageRepository) FindMessageByID(ctx context.Context, id primitive.ObjectID) (*Message, error) {
	result := r.Coll().FindOne(ctx, bson.M{"_id": id})
	if err := result.Err(); err != nil {
		return nil, err
	}
	var m *Message
	if err := result.Decode(m); err != nil {
		return nil, err
	}

	return m, nil
}

func (r *messageRepository) FindMessagesByPublicID(ctx context.Context, id primitive.ObjectID) ([]*Message, error) {
	cur, err := r.Coll().Find(ctx, bson.M{"publicId": id})
	if err != nil {
		return nil, err
	}
	var data []*Message
	if err := cur.All(ctx, &data); err != nil {
		return nil, err
	}

	return data, nil
}
func (r *messageRepository) FindMessagesByPrivateID(ctx context.Context, id primitive.ObjectID) ([]*Message, error) {
	cur, err := r.Coll().Find(ctx, bson.M{"privateID": id})
	if err != nil {
		return nil, err
	}
	var data []*Message
	if err := cur.All(ctx, &data); err != nil {
		return nil, err
	}

	return data, nil
}

// func (r *messageRepository) FindMessageByID(ctx context.Context, id primitive.ObjectID) (*Messsage, error) {
// 	result := r.Coll().FindOne(ctx, bson.M{"_id": id})
// 	if err := result.Err(); err != nil {
// 		return nil, err
// 	}
// 	var m Message
// 	if err := result.Decode(&m); err != nil {
// 		return nil, err
// 	}
	
// 	return &m, nil
// }

func (r *messageRepository) DeleteMsg(ctx context.Context, id primitive.ObjectID) error {
	result, err := r.Coll().UpdateByID(ctx, id, bson.M{"deleted": true})
	if err != nil {
		return err
	}
	_ = result

	return nil
}

func NewMessageRepository(c appctx.Container) MessageRepository {
	return &messageRepository{
		c:          c,
		repository: NewRepository(c),
	}
}

var _ MessageRepository = (*messageRepository)(nil)

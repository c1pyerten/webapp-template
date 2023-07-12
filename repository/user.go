package repository

import (
	"c1pherten/yet-webapp2/appctx"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const usersColl = "users"

var (
	ErrDupRecords = errors.New("duplicated")
	ErrNoSuchUser = errors.New("no such user")
)

type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name"`
	Nickname string             `bson:"nickname"`
	Password string             `bson:"password"`
	CreateAt time.Time          `bson:"createAt"`
}

type UserRepository interface {
	GetUser(context.Context, primitive.ObjectID) (*User, error)
	GetUserByName(context.Context, string, string) (*User, error)
	CreateNewUser(context.Context, User) error
}

type userRepository struct {
	// db *mongo.Database
	// client *mongo.Client
	c appctx.Container
	repository
}

func (r *userRepository) Coll() *mongo.Collection {
	return r.db.Collection(usersColl)
}

func (r *userRepository) GetUser(ctx context.Context, id primitive.ObjectID) (*User, error) {
	return nil, nil
}
func (r *userRepository) CreateNewUser(ctx context.Context, u User) error {
	sr := r.Coll().FindOne(ctx, bson.M{"name": u.Name})
	if sr.Err() != mongo.ErrNoDocuments {
		return ErrDupRecords
	}

	result, err := r.Coll().InsertOne(ctx, u)
	if err != nil {
		return err
	}
	_ = result

	return nil
}

func (r *userRepository) GetUserByName(ctx context.Context, name, password string) (*User, error) {
	n, err := r.Coll().CountDocuments(ctx, bson.M{"name": name})
	if err != nil {
		return nil, err
	}
	if n == 0 {
		return nil, ErrNoSuchUser
	}

	singleResult := r.Coll().FindOne(ctx, bson.M{"name": name, "password": password})
	if err := singleResult.Err(); err != nil {
		// err = errors.Join(err, ErrNoSuchUser)
		if err == mongo.ErrNoDocuments {
			err = ErrNoSuchUser
		}
		return nil, err
	}
	var u User
	if err := singleResult.Decode(&u); err != nil {
		return nil, err
	}

	return &u, nil
}

func NewUserRepository(c appctx.Container) UserRepository {
	return &userRepository{
		c:          c,
		repository: NewRepository(c),
	}
}

var _ UserRepository = (*userRepository)(nil)

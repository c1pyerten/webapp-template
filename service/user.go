package service

import (
	"c1pherten/yet-webapp2/appctx"
	"c1pherten/yet-webapp2/dto"
	"c1pherten/yet-webapp2/middleware"
	"c1pherten/yet-webapp2/repository"
	"context"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	c        appctx.Container
	userRepo repository.UserRepository
}

func (s *UserService) GetUserByID(id string) (*repository.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	u, err := s.userRepo.GetUser(ctx, oid)
	if err != nil {
		return nil, err
	}

	return u, err
}

func (s *UserService) CreateNewUser(user dto.CreateUser) (repository.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rUser := repository.User{
		ID: primitive.NewObjectID(),
		Name:     user.Name,
		Nickname: user.Nickname,
		Password: user.Password,
	}

	if err := s.userRepo.CreateNewUser(ctx, rUser); err != nil {
		return repository.User{}, err
	}
	

	// if err := s.userRepo.NewUser(ctx, repository.User{
	// 	ID:       id,
	// 	Name:     user.Name,
	// 	Nickname: user.Nickname,
	// }); err != nil {
	// 	return err
	// }

	return rUser, nil
}

// todo: f
func (s *UserService) GetUserLists() error {
	return nil
}

func (s *UserService) Login(l dto.Login) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	u, err := s.userRepo.GetUserByName(ctx, l.Name, l.Password)
	if err != nil {
		return "", err
	}
	tokenString, err := middleware.Sign(u)
	if err != nil {
		return "", err
	}

	return tokenString, err
}

func (s *UserService) QrcodeID() (string, error) {
	u := uuid.New()
	return u.String(), nil
}

func NewUserService(c appctx.Container) *UserService {
	repo := repository.NewUserRepository(c)
	return &UserService{
		c:        c,
		userRepo: repo,
	}

}

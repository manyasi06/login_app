package services

import (
	"context"
	"login_app/internal/models"
	"login_app/internal/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, user models.User) error
	DeleteUser(ctx context.Context, userId string) error
}

type UserServiceImpl struct {
	userRepository repository.UserRepostiory
}

func (u UserServiceImpl) CreateUser(ctx context.Context, user models.User) error {
	err := u.userRepository.CreateUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (u UserServiceImpl) DeleteUser(ctx context.Context, userId string) error {
	err := u.userRepository.DeleteUser(ctx, userId)
	if err != nil {
		return err
	}
	return nil
}

func NewUserService(userRepo repository.UserRepostiory) UserService {
	return UserServiceImpl{
		userRepository: userRepo,
	}
}

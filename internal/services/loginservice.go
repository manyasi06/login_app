package services

import (
	"context"
	"login_app/internal/repository"
)

type LoginService interface {
	Validate(ctx context.Context, token string) error
	Login(ctx context.Context, username, password string) (string, error)
}

type LoginServiceImpl struct {
	LoginRepo repository.LoginRepository
}

func (l LoginServiceImpl) Validate(ctx context.Context, token string) error {
	err := l.LoginRepo.Validate(ctx, token)
	if err != nil {
		return err
	}
	return nil
}

func (l LoginServiceImpl) Login(ctx context.Context, username, password string) (string, error) {
	token, err := l.LoginRepo.Login(ctx, username, password)
	if err != nil {
		return "", err
	}
	return token, nil
}

func NewLoginService(LoginRepo repository.LoginRepository) LoginService {
	return LoginServiceImpl{
		LoginRepo: LoginRepo,
	}
}

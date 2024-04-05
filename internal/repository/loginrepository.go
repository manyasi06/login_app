package repository

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"login_app/internal/config"
	"login_app/internal/models"
	"login_app/internal/util"
)

type LoginRepository interface {
	Validate(ctx context.Context, token string) error
	Login(ctx context.Context, username, password string) (string, error)
}

type LoginRepositoryImpl struct {
	Db *mongo.Database
}

func (l LoginRepositoryImpl) Login(ctx context.Context, username, password string) (string, error) {
	var result models.User
	err := l.Db.Collection("user").FindOne(ctx, bson.D{{"username", username}}).Decode(&result)
	if err != nil {
		return "", err
	}
	secretKey := config.EnvConfigs.SECRET_KEY
	decrypt, err := util.DecryptPassword(result.Password, secretKey)
	if err != nil {
		return "", err
	}
	if decrypt != password {
		return "", errors.New("Unauthorized")
	}

	token, err := util.GenerateJwtRSA(username)
	if err != nil {
		return "", err
	}
	return token, nil

}

func (l LoginRepositoryImpl) Validate(ctx context.Context, token string) error {
	_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodECDSA)
		if !ok {
			return "", errors.New("Unauthorized")
		}
		return "", nil
	})
	if err != nil {
		return err
	}
	return nil
}

func NewLoginRepository(db *mongo.Database) LoginRepository {
	return LoginRepositoryImpl{
		Db: db,
	}
}

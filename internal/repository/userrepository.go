package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"login_app/internal/config"
	"login_app/internal/models"
	"login_app/internal/util"
	"time"
)

type UserRepostiory interface {
	CreateUser(ctx context.Context, user models.User) error
	DeleteUser(ctx context.Context, id string) error
}

type UserRepositoryImpl struct {
	client *mongo.Database
}

func (u UserRepositoryImpl) CreateUser(ctx context.Context, user models.User) error {
	newPassword, err := util.Encrypt(user.Password, config.EnvConfigs.SECRET_KEY)
	user.Password = newPassword
	user.CreatedAt, err = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	one, err := u.client.Collection("user").InsertOne(ctx, user)
	if err != nil {
		return err
	}
	fmt.Println(fmt.Sprintf("Successfully inserted the following id: %s", one.InsertedID))
	return nil
}

func (u UserRepositoryImpl) DeleteUser(ctx context.Context, id string) error {
	_, err := u.client.Collection("user").UpdateByID(ctx, bson.D{{"_id", id}}, bson.D{{"isActive", "false"}})
	if err != nil {
		return err
	}
	return nil
}

func NewUserRepository(client *mongo.Database) UserRepostiory {
	return UserRepositoryImpl{
		client: client,
	}
}

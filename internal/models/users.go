package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username  string             `bson:"username" json:"username"`
	Password  string             `bson:"password" json:"password,omitempty"`
	CreatedAt time.Time          `bson:"created_At" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_At" json:"updated_at"`
	IsActive  bool               `bson:"is_active" json:"is_active,omitempty"`
}

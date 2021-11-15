package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	Id         primitive.ObjectID `json:"id" bson:"_id"`
	Title      string             `json:"title" bson:"title, omitempty"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
}

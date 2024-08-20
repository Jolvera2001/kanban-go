package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Column struct {
	ID        *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string              `json:"name" bson:"name"`
	CreatedAt time.Time           `json:"createdAt" bson:"createdAt"`
}

type ColumnDto struct {
	Name string `json:"name"`
}

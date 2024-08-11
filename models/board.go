package models

import "time"

type Board struct {
	ID        *string   `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string    `json:"name" bson:"name"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	Columns   []Column  `json:"columns,omitempty" bson:"columns,omitempty"`
}

type BoardDto struct {
	Name string `json:"name"`
}

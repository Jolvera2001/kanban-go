package services

import (
	"kanban-go/internal/database"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	columnCollectionName = "Columns"
)

type ColumnService struct {
	Collection *mongo.Collection
}

func NewColumnService() *ColumnService {
	if database.MongoClient == nil {
		panic("MongoClient isn't initialized for ColumnService to use!")
	}
	collection := database.GetCollection(database.DbName, columnCollectionName)
	return &ColumnService{Collection: collection}
}



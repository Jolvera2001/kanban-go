package services

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"kanban-go/internal/database"
	"kanban-go/internal/models"
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

func (s *ColumnService) CreateColumn(ColumnDto models.Column) (primitive.ObjectID, error) {

}

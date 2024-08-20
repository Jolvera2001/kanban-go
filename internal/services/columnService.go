package services

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"kanban-go/internal/database"
	"kanban-go/internal/models"
)

const (
	columnCollectionName = "Columns"
)

var notImplementedError = errors.New("not Implemented")

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
	return primitive.NewObjectID(), notImplementedError
}

func (s *ColumnService) UpdateColumn(ColumnDto models.Column) error {
	return notImplementedError
}

func (s *ColumnService) DeleteColumn(id primitive.ObjectID) error {
	return notImplementedError
}

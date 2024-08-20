package services

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"kanban-go/internal/database"
	"kanban-go/internal/models"
	"time"
)

const (
	columnCollectionName = "Boards"
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

func (s *ColumnService) CreateColumn(boardId primitive.ObjectID, ColumnDto models.ColumnDto) (primitive.ObjectID, error) {
	createdColumn, err := createColumn(ColumnDto)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	filter := bson.M{
		"_id": boardId,
	}
	update := bson.M{
		"$push": bson.M{
			"columns": createdColumn,
		},
	}

	res, err := s.Collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return res.UpsertedID.(primitive.ObjectID), nil
}

func (s *ColumnService) UpdateColumn(boardId primitive.ObjectID, Column models.Column) error {
	filter := bson.M{"_id": boardId, "columns._id": Column.ID}
	update := bson.M{"$set": bson.M{"columns.$": Column}}

	_, err := s.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (s *ColumnService) DeleteColumn(boardId primitive.ObjectID, columnId primitive.ObjectID) error {
	filter := bson.M{"_id": boardId, "columns._id": columnId}
	_, err := s.Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}

func createColumn(ColumnDto models.ColumnDto) (models.Column, error) {
	var column = models.Column{
		Name:      ColumnDto.Name,
		CreatedAt: time.Now(),
	}

	return column, nil
}

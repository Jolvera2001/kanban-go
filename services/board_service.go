package services

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"kanban-go/database"
	"kanban-go/models"
)

const (
	dbName         = "Main"
	collectionName = " Boards"
)

var collection = database.GetCollection(dbName, collectionName)

type MongoBoardsService struct {
	collection *mongo.Collection
}

func NewMongoBoardsService() *MongoBoardsService {
	return &MongoBoardsService{collection: collection}
}

func (s *MongoBoardsService) CreateBoard(board models.Board) error {
	_, err := s.collection.InsertOne(context.TODO(), board)
	return err
}

func (s *MongoBoardsService) GetBoards() ([]models.Board, error) {
	var boards []models.Board
	cursor, err := s.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	defer func() {
		if closeErr := cursor.Close(context.TODO()); closeErr != nil {
			err = closeErr
		}
	}()

	boards, err = parseCursor(&boards, cursor)
	if err != nil {
		return nil, err
	}

	return boards, nil
}

func (s *MongoBoardsService) GetBoardById(id string) (*models.Board, error) {
	var board models.Board
	var filter = bson.M{"id": id}

	err := s.collection.FindOne(context.TODO(), filter).Decode(&board)
	return &board, err
}

func (s *MongoBoardsService) UpdateBoard(Board *models.Board) error {
	_, err := s.collection.UpdateByID(context.TODO(), Board.ID, Board)
	if err != nil {
		return err
	}

	return nil
}

func (s *MongoBoardsService) DeleteBoard(id string) error {
	var filter = bson.M{"_id": id}
	_, err := s.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil
}

func parseCursor(boards *[]models.Board, cursor *mongo.Cursor) ([]models.Board, error) {
	for cursor.Next(context.TODO()) {
		var board models.Board
		if err := cursor.Decode(&board); err != nil {
			return nil, err
		}
		*boards = append(*boards, board)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return *boards, nil
}

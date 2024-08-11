package services

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"kanban-go/database"
	"kanban-go/models"
	"time"
)

const (
	dbName         = "Main"
	collectionName = " Boards"
)

type MongoBoardsService struct {
	collection *mongo.Collection
}

func NewMongoBoardsService() *MongoBoardsService {
	if database.MongoClient == nil {
		panic("MongoClient isn't initialized for BoardService to use!")
	}
	collection := database.GetCollection(dbName, collectionName)
	return &MongoBoardsService{collection: collection}
}

func (s *MongoBoardsService) CreateBoard(BoardDto models.BoardDto) error {
	board, err := createDefaultBoard(BoardDto)
	_, err = s.collection.InsertOne(context.TODO(), board)
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
	var filter = bson.M{"_id": id}

	err := s.collection.FindOne(context.TODO(), filter).Decode(&board)
	return &board, err
}

func (s *MongoBoardsService) UpdateBoard(Board models.Board) error {
	var filter = bson.M{"_id": Board.ID}
	_, err := s.collection.UpdateOne(context.TODO(), filter, Board)
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

func createDefaultBoard(BoardDto models.BoardDto) (models.Board, error) {
	var board models.Board = models.Board{
		Name: BoardDto.Name,
		Columns: []models.Column{
			{}, // TODO: decide on what should be added for default columns
		},
		CreatedAt: time.Now(),
	}

	return board, nil
}

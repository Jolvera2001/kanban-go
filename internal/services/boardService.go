package services

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"kanban-go/internal/database"
	models2 "kanban-go/internal/models"
	"time"
)

const (
	collectionName = " Boards"
)

type BoardService struct {
	Collection *mongo.Collection
}

func NewBoardsService() *BoardService {
	if database.MongoClient == nil {
		panic("MongoClient isn't initialized for IBoardService to use!")
	}
	collection := database.GetCollection(database.DbName, collectionName)
	return &BoardService{Collection: collection}
}

func (s *BoardService) CreateBoard(BoardDto models2.BoardDto) (primitive.ObjectID, error) {
	board, err := createDefaultBoard(BoardDto)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	res, err := s.Collection.InsertOne(context.TODO(), board)
	if res == nil {
		return primitive.ObjectID{}, err
	}
	return res.InsertedID.(primitive.ObjectID), err
}

func (s *BoardService) GetBoards() ([]models2.Board, error) {
	var boards []models2.Board
	cursor, err := s.Collection.Find(context.TODO(), bson.M{})
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

func (s *BoardService) GetBoardById(id primitive.ObjectID) (models2.Board, error) {
	var board models2.Board
	var filter = bson.M{"_id": id}

	err := s.Collection.FindOne(context.TODO(), filter).Decode(&board)
	return board, err
}

func (s *BoardService) UpdateBoard(Board models2.Board) error {
	var filter = bson.M{"_id": Board.ID}
	updateData, err := bson.Marshal(Board)
	if err != nil {
		return err
	}

	var update bson.M
	if err = bson.Unmarshal(updateData, &update); err != nil {
		return err
	}

	_, err = s.Collection.UpdateOne(context.TODO(), filter, bson.M{"$set": update})
	if err != nil {
		return err
	}

	return nil
}

func (s *BoardService) DeleteBoard(id primitive.ObjectID) error {
	var filter = bson.M{"_id": id}
	_, err := s.Collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil
}

func parseCursor(boards *[]models2.Board, cursor *mongo.Cursor) ([]models2.Board, error) {
	for cursor.Next(context.TODO()) {
		var board models2.Board
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

func createDefaultBoard(BoardDto models2.BoardDto) (models2.Board, error) {
	var board models2.Board = models2.Board{
		Name: BoardDto.Name,
		Columns: []models2.Column{
			{}, // TODO: decide on what should be added for default columns
		},
		CreatedAt: time.Now(),
	}

	return board, nil
}

package services

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"kanban-go/database"
	"kanban-go/models"
	"time"
)

const (
	dbName         = "Main"
	collectionName = " Boards"
)

type BoardService struct {
	collection *mongo.Collection
}

func NewBoardsService() *BoardService {
	if database.MongoClient == nil {
		panic("MongoClient isn't initialized for IBoardService to use!")
	}
	collection := database.GetCollection(dbName, collectionName)
	return &BoardService{collection: collection}
}

func (s *BoardService) CreateBoard(BoardDto models.BoardDto) (primitive.ObjectID, error) {
	board, err := createDefaultBoard(BoardDto)
	res, err := s.collection.InsertOne(context.TODO(), board)
	if res == nil {
		return primitive.ObjectID{}, err
	}
	return res.InsertedID.(primitive.ObjectID), err
}

func (s *BoardService) GetBoards() ([]models.Board, error) {
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

func (s *BoardService) GetBoardById(id primitive.ObjectID) (models.Board, error) {
	var board models.Board
	var filter = bson.M{"_id": id}

	err := s.collection.FindOne(context.TODO(), filter).Decode(&board)
	return board, err
}

func (s *BoardService) UpdateBoard(Board models.Board) error {
	var filter = bson.M{"_id": Board.ID}
	updateData, err := bson.Marshal(Board)
	if err != nil {
		return err
	}

	var update bson.M
	if err = bson.Unmarshal(updateData, &update); err != nil {
		return err
	}

	_, err = s.collection.UpdateOne(context.TODO(), filter, bson.M{"$set": update})
	if err != nil {
		return err
	}

	return nil
}

func (s *BoardService) DeleteBoard(id primitive.ObjectID) error {
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

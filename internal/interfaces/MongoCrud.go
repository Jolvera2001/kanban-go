package interfaces

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"kanban-go/internal/models"
)

type IBoardService interface {
	CreateBoard(BoardDto models.BoardDto) (primitive.ObjectID, error)
	GetBoards() ([]models.Board, error)
	GetBoardById(id primitive.ObjectID) (models.Board, error)
	UpdateBoard(Board models.Board) error
	DeleteBoard(id primitive.ObjectID) error
}

type MongoCrud[TInput any, TOutput any] interface {
	Create(entity TInput) (TOutput, error)
	ReadById(id primitive.ObjectID) (TOutput, error)
	Update(entity TInput) error
	Delete(id primitive.ObjectID) error
}

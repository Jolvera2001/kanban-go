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

type IColumnService interface {
	CreateColumn(boardId primitive.ObjectID, ColumnDto models.ColumnDto) (primitive.ObjectID, error)
	UpdateColumn(boardId primitive.ObjectID, Column models.Column) error
	DeleteColumn(boardId primitive.ObjectID, columnId primitive.ObjectID) error
}

package services

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"kanban-go/models"
)

type BoardService interface {
	CreateBoard(BoardDto models.Board) error
	GetBoards() ([]models.Board, error)
	GetBoardById(id primitive.ObjectID) error
	UpdateBoard(Board models.Board) error
	DeleteBoard(id primitive.ObjectID) error
}

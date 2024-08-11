package services

import (
	"kanban-go/models"
)

type BoardService interface {
	CreateBoard(BoardDto models.Board) error
	GetBoards() ([]models.Board, error)
	GetBoardById(id string) error
	UpdateBoard(Board models.Board) error
	DeleteBoard(id string) error
}

package services

import (
	"kanban-go/models"
)

type BoardService interface {
	CreateBoard(Board *models.Board) error
	GetBoards() ([]models.Board, error)
	GetBoardById(id string) error
	UpdateBoard(Board *models.Board) error
	DeleteBoard(id string) error
}

package Mocks

import (
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"kanban-go/models"
)

type MockBoardService struct {
	mock.Mock
}

func (m *MockBoardService) CreateBoard(BoardDto models.BoardDto) (primitive.ObjectID, error) {
	args := m.Called(BoardDto)
	return args.Get(0).(primitive.ObjectID), args.Error(1)
}

func (m *MockBoardService) GetBoards() ([]models.Board, error) {
	args := m.Called()
	return args.Get(0).([]models.Board), args.Error(1)
}

func (m *MockBoardService) GetBoardById(id primitive.ObjectID) (models.Board, error) {
	args := m.Called(id)
	return args.Get(0).(models.Board), args.Error(1)
}

func (m *MockBoardService) UpdateBoard(Board models.Board) error {
	args := m.Called(Board)
	return args.Error(0)
}

func (m *MockBoardService) DeleteBoard(id primitive.ObjectID) error {
	args := m.Called(id)
	return args.Error(0)
}

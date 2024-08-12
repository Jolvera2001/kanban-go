package services

import (
	"kanban-go/mocks"
	"testing"
)

func TestCreateBoard(t *testing.T) {
	mockCollection := new(mocks.MockMongoCollection)
	service := NewMongoBoardsService(mockCollection)
}

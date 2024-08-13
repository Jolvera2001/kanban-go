package serviceTests

import (
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"kanban-go/internal/database"
	"kanban-go/internal/models"
	"kanban-go/internal/services"
	"os"
	"testing"
)

func NewTestBoardsService() *services.BoardService {
	connStr := os.Getenv("KB_DB")
	err := database.ConnectToMongoDB(connStr)
	if err != nil {
		return nil
	}

	if database.MongoClient == nil {
		panic("MongoClient isn't initialized for IBoardService to use!")
	}
	collection := database.GetCollection("Test", "BoardsTest")
	return &services.BoardService{Collection: collection}
}

func TestBoardService_CreateBoardPass(t *testing.T) {
	service := NewTestBoardsService()
	newBoard := models.BoardDto{Name: "Integration Test Board! Create Test"}

	res, err := service.CreateBoard(newBoard)

	assert.NoError(t, err, "expected no error")
	assert.IsType(t, primitive.ObjectID{}, res, "expected a primitive.ObjectID")

	// cleanup
	err = service.DeleteBoard(res)
	if err != nil {
		return
	}
}

func TestBoardService_GetBoardsPass(t *testing.T) {
	service := NewTestBoardsService()

	res, err := service.GetBoards()

	assert.NoError(t, err, "expected no error")
	assert.IsType(t, []models.Board{}, res, "expected []models.Board")
}

func TestBoardService_GetBoardByIdPass(t *testing.T) {
	service := NewTestBoardsService()
	newBoard := models.BoardDto{Name: "Integration Test Board! Get By Id"}
	id, err := service.CreateBoard(newBoard)
	if err != nil {
		return
	}

	res, err := service.GetBoardById(id)

	assert.NoError(t, err, "expected no error")
	assert.IsType(t, models.Board{}, res, "expected a models.Board")

	// cleanup
	err = service.DeleteBoard(id)
	if err != nil {
		return
	}
}

func TestBoardService_UpdateBoardPass(t *testing.T) {
	service := NewTestBoardsService()
	newBoard := models.BoardDto{Name: "Integration Test Board! Update Test"}
	res, err := service.CreateBoard(newBoard)
	if err != nil {
		return
	}
	currBoard, err := service.GetBoardById(res)
	if err != nil {
		return
	}
	updateBoard := currBoard
	updateBoard.Name = "Integration Test Board: Updated test board haha"

	err = service.UpdateBoard(updateBoard)
	check, err := service.GetBoardById(res)

	assert.NoError(t, err, "expected no error")
	assert.Equal(t, updateBoard, check, "UpdateBoard and Check should be equal")

	// cleanup
	err = service.DeleteBoard(res)
}

func TestBoardService_DeleteBoardPass(t *testing.T) {
	service := NewTestBoardsService()
	newBoard := models.BoardDto{Name: "Integration Test Board! Delete Test"}
	id, err := service.CreateBoard(newBoard)
	if err != nil {
		return
	}

	err = service.DeleteBoard(id)

	assert.NoError(t, err, "expected no error")
}

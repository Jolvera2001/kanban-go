package routeTests

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	models2 "kanban-go/internal/models"
	"kanban-go/internal/routes"
	"kanban-go/tests/Mocks"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestBoardRoutes_MethodCalling(t *testing.T) {
	router := gin.Default()
	mockService := new(Mocks.MockBoardService)
	routes.BoardRoutes(router, mockService)

	boardDto := models2.BoardDto{Name: "Test Board"}
	id := primitive.NewObjectID()
	createdTime := time.Now()
	board := models2.Board{
		ID:        &id,
		Name:      "TestBoard",
		CreatedAt: createdTime,
		Columns:   nil,
	}

	// expectations
	mockService.On("CreateBoard", boardDto).Return(primitive.NewObjectID(), nil)
	mockService.On("GetBoards").Return([]models2.Board{board}, nil)
	mockService.On("GetBoardById", mock.Anything).Return(board, nil)
	mockService.On("UpdateBoard", mock.Anything).Return(nil)
	mockService.On("DeleteBoard", mock.Anything).Return(nil)

	req, _ := http.NewRequest("GET", "/api/v1/boards", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	boardId := board.ID.Hex()
	req, _ = http.NewRequest("GET", "/api/v1/board/"+boardId, nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	boardDtoJSON, _ := json.Marshal(boardDto)
	req, _ = http.NewRequest("POST", "/api/v1/board", bytes.NewBuffer(boardDtoJSON))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	boardJSON, _ := json.Marshal(board)
	req, _ = http.NewRequest("PUT", "/api/v1/board", bytes.NewBuffer(boardJSON))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	reqJSON, _ := json.Marshal(models2.IdRequest{ID: boardId})
	req, _ = http.NewRequest("DELETE", "/api/v1/board", bytes.NewBuffer(reqJSON))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	mockService.AssertExpectations(t)
}

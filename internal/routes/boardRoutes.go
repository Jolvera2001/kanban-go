package routes

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	models2 "kanban-go/internal/models"
	"kanban-go/internal/services"
	"net/http"
)

func BoardRoutes(r *gin.Engine, service services.IBoardService) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/boards", func(c *gin.Context) {
			boards, err := service.GetBoards()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"Boards": boards})
		})

		v1.GET("/board/:id", func(c *gin.Context) {
			reqParam := c.Param("id")

			boardId, err := primitive.ObjectIDFromHex(reqParam)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			board, err := service.GetBoardById(boardId)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{"Board": board})
		})

		v1.POST("/board", func(c *gin.Context) {
			var boardDto models2.BoardDto
			if err := c.ShouldBindJSON(&boardDto); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			res, err := service.CreateBoard(boardDto)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			c.JSON(http.StatusOK, gin.H{"Board": res})
		})

		v1.PUT("/board", func(c *gin.Context) {
			var board models2.Board
			if err := c.ShouldBindJSON(&board); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			if err := service.UpdateBoard(board); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"status": "Board Updated"})
		})

		v1.DELETE("/board", func(c *gin.Context) {
			var req models2.IdRequest
			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			boardId, err := primitive.ObjectIDFromHex(req.ID)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			if err := service.DeleteBoard(boardId); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"status": "Board Deleted"})
		})
	}
}

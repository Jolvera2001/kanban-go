package routes

import (
	"github.com/gin-gonic/gin"
	"kanban-go/models"
	"kanban-go/services"
	"net/http"
)

func BoardRoutes(r *gin.Engine) {
	service := services.NewMongoBoardsService()

	r.GET("/boards", func(c *gin.Context) {
		boards, err := service.GetBoards()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"Boards": boards})
	})

	r.POST("/board", func(c *gin.Context) {
		var board models.Board
		if err := c.ShouldBindJSON(&board); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := service.CreateBoard(board); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "Board Created"})
	})
}

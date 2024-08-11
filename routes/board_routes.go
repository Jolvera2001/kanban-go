package routes

import (
	"github.com/gin-gonic/gin"
	"kanban-go/models"
	"kanban-go/services"
	"net/http"
)

func BoardRoutes(r *gin.Engine) {
	service := services.NewMongoBoardsService()

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

		v1.POST("/board", func(c *gin.Context) {
			var boardDto models.BoardDto
			if err := c.ShouldBindJSON(&boardDto); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			if err := service.CreateBoard(boardDto); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"status": "Board Created"})
		})
	}
}

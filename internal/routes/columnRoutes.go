package routes

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"kanban-go/internal/models"
	"kanban-go/internal/services"
	"net/http"
)

func ColumnRoutes(r *gin.Engine, service services.ColumnService) {
	v1 := r.Group("/api/v1")
	{
		v1.POST("/column", func(c *gin.Context) {
			var column columnRequest
			if err := c.ShouldBindJSON(&column); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			res, err := service.CreateColumn(column.BoardId, column.Column)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			c.JSON(http.StatusCreated, res)
		})

		v1.PUT("/column", func(c *gin.Context) {
			var column columnFullRequest
			if err := c.ShouldBindJSON(&column); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			err := service.UpdateColumn(column.BoardId, column.Column)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		})

		v1.DELETE("/column", func(c *gin.Context) {
			var column columnFullRequest
			if err := c.ShouldBindJSON(&column); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}

			err := service.DeleteColumn(column.BoardId, *column.Column.ID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		})
	}
}

type columnRequest struct {
	BoardId primitive.ObjectID `json:"BoardId"`
	Column  models.ColumnDto   `json:"Column"`
}

type columnFullRequest struct {
	BoardId primitive.ObjectID `json:"BoardId"`
	Column  models.Column      `json:"Column"`
}

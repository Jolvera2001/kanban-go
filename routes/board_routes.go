package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func boardRoutes(r *gin.Engine) {
	r.GET("/boards", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Boards": "Some boards here idk",
		})
	})

	r.POST("/boards", func(c *gin.Context) {
		// define structs for validation
		c.JSON(http.StatusOK, gin.H{"Status": "Did post request"})
	})
}

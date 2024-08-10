package main

import (
	"github.com/gin-gonic/gin"
	"kanban-go/database"
	"log"
	"net/http"
	"os"
)

func routes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/login", func(c *gin.Context) {
		var json struct {
			User     string `json:"user" binding:"required"`
			Password string `json:"password" binding:"required"`
		}

		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		if json.User != "admin" || json.Password != "password" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"Status": "You are logged in!"})
	})
}

func main() {
	// env variables
	dbHost := os.Getenv("KB_DB")
	if dbHost == "" {
		log.Fatal("KB_DB environment variable not set")
	}

	// routes
	router := gin.Default()
	routes(router)

	// init db connection
	database.ConnectToMongoDB(dbHost)

	err := router.Run(":8080")
	if err != nil {
		return
	}
}

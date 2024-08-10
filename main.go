package main

import (
	"github.com/gin-gonic/gin"
	"kanban-go/database"
	"kanban-go/routes"
	"log"
	"os"
)

func main() {
	// env variables
	dbHost := os.Getenv("KB_DB")
	if dbHost == "" {
		log.Fatal("KB_DB environment variable not set")
	}

	// routes
	router := gin.Default()
	routes.BoardRoutes(router)

	// init db connection
	database.ConnectToMongoDB(dbHost)

	err := router.Run(":8080")
	if err != nil {
		return
	}
}

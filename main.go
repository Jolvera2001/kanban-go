package main

import (
	"github.com/gin-gonic/gin"
	"kanban-go/database"
	"kanban-go/routes"
	"kanban-go/services"
	"log"
	"os"
)

func main() {
	router := gin.Default()
	connStr := os.Getenv("KB_DB")
	if connStr == "" {
		log.Fatal("KB_DB environment variable not set")
	}

	// set trusted proxies
	//err := router.SetTrustedProxies([]string{"192.168.1.1", "10.0.0.0/24"})
	//if err != nil {
	//	return
	//}

	// init db connection
	err := database.ConnectToMongoDB(connStr)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	boardService := services.NewBoardsService()

	// routes
	routes.BoardRoutes(router, boardService)

	if err = router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

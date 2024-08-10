package services

import (
	"go.mongodb.org/mongo-driver/mongo"
	"kanban-go/database"
)

const (
	dbName         = "Main"
	collectionName = " Boards"
)

var collection *mongo.Collection = database.CreateCollection(dbName, collectionName)

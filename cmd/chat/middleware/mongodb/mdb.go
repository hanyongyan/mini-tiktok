package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoDB manages MongoDB connection
type MongoDB struct {
	Cli *mongo.Client
	Databasename string
}

var (
	Mongo *MongoDB
)
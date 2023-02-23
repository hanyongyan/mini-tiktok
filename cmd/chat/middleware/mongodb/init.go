package mongodb

import (
	"context"
	"fmt"
	"mini_tiktok/pkg/configs/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Init initializes mongo database
func Init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	conf := config.GlobalConfigs.MongoDbConfig
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf(
		"mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority",
		conf.Username,
		conf.Password,
		conf.Addr,
	)))
	if err != nil {
		panic(err)
	}
	Mongo = &MongoDB{Cli: client, Databasename: conf.DbName}
}

// Close the existing connection
func (db *MongoDB) Close() {
	if db.Cli != nil {
		db.Cli.Disconnect(context.TODO())
	}
}

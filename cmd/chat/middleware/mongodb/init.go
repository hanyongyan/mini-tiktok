package mongodb

import (
	"gopkg.in/mgo.v2"
	"mini_tiktok/pkg/configs/config"
	"time"
)

// Init initializes mongo database
func Init() {
	db := &MongoDB{}
	conf := config.GlobalConfigs.MongoDbConfig
	db.Databasename = conf.DbName

	// DialInfo holds options for establishing a session with a MongoDB cluster.
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{conf.Addr}, // Get HOST + PORT
		Timeout:  60 * time.Second,
		Database: db.Databasename, // Database name
		Username: conf.Username,   // Username
		Password: conf.Password,   // Password
	}

	// Create a session which maintains a pool of socket connections
	// to the DB MongoDB database.
	var err error
	db.DbSession, err = mgo.DialWithInfo(dialInfo)

	if err != nil {
		panic(err)
	}
	Mongo = db
}

// Close the existing connection
func (db *MongoDB) Close() {
	if db.DbSession != nil {
		db.DbSession.Close()
	}
}

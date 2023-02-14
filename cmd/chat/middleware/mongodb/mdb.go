package mongodb

import (
	mgo "gopkg.in/mgo.v2"
)

// MongoDB manages MongoDB connection
type MongoDB struct {
	DbSession  *mgo.Session
	Databasename string
}

var (
	Mongo *MongoDB
)
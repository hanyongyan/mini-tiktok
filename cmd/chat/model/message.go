package model

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

// Message information
type Message struct {
	ID         bson.ObjectId `bson:"_id" json:"id"`
	ChatKey    string        `bson:"chat_key" json:"chat_key"`
	Content    string        `bson:"content" json:"content"`
	CreateTime time.Time `bson:"create_time" json:"create_time"`
}

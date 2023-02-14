package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
	"mini_tiktok/cmd/chat/middleware/mongodb"
	"mini_tiktok/cmd/chat/model"
	"time"
)

type MsgDao struct {
	collection *mongo.Collection
}

const (
	COLLECTION = "messages"
)

func (dao *MsgDao) initSession() {
	dbname := mongodb.Mongo.Databasename
	dao.collection = mongodb.Mongo.Cli.Database(dbname).Collection(COLLECTION)
}

func (dao *MsgDao) GetAllByChatKey(ctx context.Context, chatKey string) ([]model.Message, error) {
	dao.initSession()
	var messages []model.Message
	find, err := dao.collection.Find(ctx, bson.M{"chat_key": chatKey})
	if err != nil {
		return messages, err
	}
	err = find.All(ctx, &messages)
	return messages, err
}

func (dao *MsgDao) Insert(ctx context.Context, chatKey, content string) error {
	dao.initSession()
	_, err := dao.collection.InsertOne(ctx, bson.M{
		"chat_key":    chatKey,
		"content":     content,
		"create_time": time.Now(),
	})
	return err
}

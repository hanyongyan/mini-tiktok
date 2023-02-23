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

func (dao *MsgDao) GetAllByChatKey(ctx context.Context, fromUserId, toUserId int64) ([]model.Message, error) {
	dao.initSession()
	var messages []model.Message
	find, err := dao.collection.Find(ctx, bson.M{"from_user_id": fromUserId, "to_user_id": toUserId})
	if err != nil {
		return messages, err
	}
	err = find.All(ctx, &messages)
	return messages, err
}

func (dao *MsgDao) Insert(ctx context.Context, fromUserId, toUserId int64, content string) error {
	dao.initSession()
	_, err := dao.collection.InsertOne(ctx, bson.M{
		"from_user_id": fromUserId,
		"to_user_id":   toUserId,
		"content":      content,
		"create_time":  time.Now().UnixNano(),
	})
	return err
}

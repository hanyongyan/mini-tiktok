package dao

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"mini_tiktok/cmd/chat/middleware/mongodb"
	"mini_tiktok/cmd/chat/model"
	"time"
)

type MsgDao struct {
	session    *mgo.Session
	collection *mgo.Collection
}

const (
	COLLECTION = "messages"
)

func (dao *MsgDao) initSession() {
	dao.session = mongodb.Mongo.DbSession.Copy()
	dao.collection = dao.session.DB(mongodb.Mongo.Databasename).C(COLLECTION)
}

func (dao *MsgDao) close() {
	dao.session.Close()
}

func (dao *MsgDao) GetAllByChatKey(chatKey string) ([]model.Message, error) {
	dao.initSession()
	defer dao.close()
	var messages []model.Message
	err := dao.collection.Find(bson.M{"chat_key": chatKey}).All(&messages)
	return messages, err
}

func (dao *MsgDao) Insert(chatKey, content string) error {
	dao.initSession()
	defer dao.close()
	msg := model.Message{ChatKey: chatKey, Content: content, CreateTime: time.Now()}
	return dao.collection.Insert(msg)
}

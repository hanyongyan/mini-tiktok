package model

// Message information
type Message struct {
	ID         string `bson:"_id" json:"id"`
	FromUserId int64  `bson:"from_user_id" json:"from_user_id"`
	ToUserId   int64  `bson:"to_user_id" json:"to_user_id"`
	Content    string `bson:"content" json:"content"`
	CreateTime int64  `bson:"create_time" json:"create_time"`
}

// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameTFollow = "t_follow"

// TFollow mapped from table <t_follow>
type TFollow struct {
	ID         int64 `gorm:"column:id;type:int(10) unsigned;primaryKey;autoIncrement:true" json:"id"` // 主键id
	UserID     int64 `gorm:"column:user_id;type:int(11);not null" json:"user_id"`                     // 用户id
	FollowerID int64 `gorm:"column:follower_id;type:int(11);not null" json:"follower_id"`             // 关注者id
}

// TableName TFollow's table name
func (*TFollow) TableName() string {
	return TableNameTFollow
}
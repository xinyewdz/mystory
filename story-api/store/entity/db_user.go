package entity

import "time"

const (
	USER_TYPE_ADMIN = "admin"
	USER_TYPE_USER  = "user"
)

type DBUser struct {
	Id         string    `json:"id" bson:"_id" db:"id"`
	Openid     string    `json:"openid" bson:"openId" db:"open_id"`
	Name       string    `json:"name" bson:"name" db:"name"`
	Gender     string    `json:"gender" bson:"gender" db:"gender"`
	Phone      string    `json:"phone" bson:"phone" db:"phone"`
	Password   string    `json:"password" bson:"password" db:"password"`
	AvatarUrl  string    `json:"avatarUrl" bson:"avatarUrl"`
	Type       string    `json:"type" bson:"type" db:"type"`
	Province   string    `json:"province" bson:"province"`
	City       string    `json:"city" bson:"city"`
	CreateTime time.Time `json:"create_time" bson:"createTime" db:"create_time"`
}

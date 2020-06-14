package entity

import "time"

type DBUser struct {
	Id string `json:"id" bson:"_id" db:"id"`
	Openid string `json:"openid" bson:"openId" db:"open_id"`
	Name string `json:"name" bson:"name" db:"name"`
	Gender string `json:"gender" bson:"gender" db:"gender"`
	Phone string `json:"phone" bson:"phone" db:"phone"`
	Password string `json:"password" bson:"password" db:"password"`
	Type string `json:"type" bson:"type" db:"type"`
	CreateTime time.Time `json:"create_time" bson:"createTime" db:"create_time"`
	
}

package entity

import "time"

type DBUser struct {
	Id int64 `json:"id"`
	Openid string `json:"openid"`
	Name string `json:"name"`
	Gender string `json:"gender"`
	Phone string `json:"phone"`
	Password string `json:"password"`
	Type string `json:"type"`
	CreateTime time.Time `json:"create_time"`
	
}

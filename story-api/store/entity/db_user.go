package entity

import "time"

type DBUser struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Phone string `json:"phone"`
	Password string `json:"password"`
	WxId string `json:"wx_id"`
	Type string `json:"type"`
	CreateTime time.Time `json:"create_time"`
	
}

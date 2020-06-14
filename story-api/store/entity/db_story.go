package entity

import "time"

type DBStory struct {
	Id string `json:"id" bson:"_id,omitempty" db:"id"`
	Name string `json:"name" bson:"name" db:"name"`
	AudioUrl string `json:"audioUrl" bson:"audioUrl" db:"audio_url"`
	ImageUrl string `json:"imageUrl" bson:"imageUrl" db:"image_url"`
	CreateTime time.Time `json:"createTime" bson:"createTime" db:"create_time"`
}

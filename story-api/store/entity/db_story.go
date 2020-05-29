package entity

import "time"

type DBStory struct {
	Id string `json:"id" bson:"_id,omitempty"`
	Name string `json:"name" bson:"name"`
	ImageUrl string `json:"imageUrl" bson:"imageUrl"`
	AudioUrl string `json:"audioUrl" bson:"audioUrl"`
	CreateTime time.Time `json:"create_time" bson:"createTime"`
}

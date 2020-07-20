package entity

import "time"

type DBStoryPlayDetail struct {
	Id         string    `json:"id" bson:"_id"`
	StoryId    string    `json:"storyId" bson:"storyId"`
	UserId     string    `json:"userId" bson:"userId"`
	CreateTime time.Time `json:"createTime" bson:"createTime"`
}

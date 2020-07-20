package entity

import "time"

type DBStoryFavorite struct {
	Id         string    `json:"id" bson:"_id"`
	UserId     string    `json:"userId" bson:"userId"`
	StoryId    string    `json:"storyId" bson:"storyId"`
	CreateTime time.Time `json:"createTime" bson:"createTime"`
}

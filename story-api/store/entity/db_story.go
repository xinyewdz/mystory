package entity

import "time"

type DBStory struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	ImageUrl string `json:"imageUrl"`
	AudioUrl string `json:"audioUrl"`
	CreateTime time.Time `json:"create_time"`
}

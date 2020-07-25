package model

type StoryPlayResp struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	ImageUrl   string `json:"imageUrl"`
	AudioUrl   string `json:"audioUrl"`
	CreateUser string `json:"createUser"`
	TotalPlay  int64  `json:"totalPlay"`
	CreateTime string `json:"createTime"`
}

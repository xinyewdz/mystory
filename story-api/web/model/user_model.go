package model

type LoginReq struct {
	Code     string `json:"code"`
	NickName string `json:"nickName"`
	Gender   string `json:"gender"`
}

type AccountResp struct {
	Id              int64  `json:"id"`
	Openid          string `json:"openid"`
	Name            string `json:"name"`
	Phone           string `json:"phone"`
	Token           string `json:"token"`
	Type            string `json:"type"`
	EnableStoryList bool   `json:"enableStoryList"`
}

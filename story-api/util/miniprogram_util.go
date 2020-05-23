package util

import (
	"encoding/json"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
)

var(
	miniprogramHost = "https://api.weixin.qq.com/sns/jscode2session"
	appId string = ""
	secret string = ""

)

func init(){
	appId = confMap["wx.appid"]
	secret = confMap["wx.secretkey"]
}

type SessionResp struct {
	ErrCode int `json:"errcode"`
	ErrMsg string `json:"errmsg"`
	OpenId string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId string `json:"unionid"`
}

func Code2Session(jsCode string)*SessionResp{
	var url string = miniprogramHost+"?appid="+appId+"&secret="+secret+"&js_code="+jsCode+"&grant_type=authorization_code"
	resp,err := http.Get(url)
	if err!=nil{
		log.Error("miniprogram code2session error",zap.Error(err))
		return nil
	}
	defer resp.Body.Close()
	data,_ := ioutil.ReadAll(resp.Body)
	sResp := &SessionResp{}
	json.Unmarshal(data,sResp)
	if sResp.ErrCode!=0{
		log.Error("miniprogram code2session fail.",zap.Int("errcode",sResp.ErrCode),zap.String("errmsg",sResp.ErrMsg))
		return nil
	}
	return sResp
}

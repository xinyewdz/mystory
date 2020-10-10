package web

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"story-api/common"
	"story-api/global"
	"story-api/store/dao/mongodao"
)

const (
	TOKEN_KEY = "s:token:"
)

var (
	mainLog            = global.MainLog
	userDao            = mongodao.NewUserDao()
	storyDao           = mongodao.NewStoryDao()
	storyPlayDetailDao = mongodao.NewStoryPlayDetailDao()
	storyFavoriteDao   = mongodao.NewStoryFavoriteDao()
)

func parseBody(reqBody *http.Request) map[string]string {
	data, _ := ioutil.ReadAll(reqBody.Body)
	reqMap := make(map[string]string)
	json.Unmarshal(data, &reqMap)
	return reqMap
}

func ResolveBody(req *http.Request, reqBody interface{}) {
	data, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(data, reqBody)
}

func WriteResp(resp http.ResponseWriter, apiResp *common.ApiResponse) {
	respData, _ := json.Marshal(apiResp)
	resp.Write(respData)
}

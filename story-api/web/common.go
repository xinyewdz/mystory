package web

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"story-api/global"
)

const(
	TOKEN_KEY = "s:token:"
)

var mainLog = global.MainLog

func parseBody(reqBody *http.Request)map[string]string{
	data,_ := ioutil.ReadAll(reqBody.Body)
	reqMap := make(map[string]string)
	json.Unmarshal(data,&reqMap)
	return reqMap
}



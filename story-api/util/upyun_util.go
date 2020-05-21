package util

import (
	"go.uber.org/zap"
	"io"
	"net/http"
	"story-api/global"
)

var(
	mainLog = global.MainLog
	downloadHost = "http://up.wenqiuqiu.com"
	nameSpace = "aaronimage/"
	host = confMap["upyun.host"]
	user = confMap["upyun.user"]
	password = confMap["upyun.password"]
)


func UpyunUpload(read io.ReadCloser,path string)string{
	uploadPath := nameSpace+path;
	client := &http.Client{}
	req, err := http.NewRequest("POST", host+uploadPath, read)
	if err!=nil {
		mainLog.Error("upyun upload error",zap.Error(err))
		return ""
	}
	req.SetBasicAuth(user, password)
	client.Do(req)
	return downloadHost+path
}
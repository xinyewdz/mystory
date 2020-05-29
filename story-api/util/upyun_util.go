package util

import (
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"net/http"
	"story-api/util/config"
)

var(
	downloadHost = "http://up.wenqiuqiu.com"
	nameSpace = "/aaronimage/"
	host string
	user string
	password string
)

func init()  {
	host = config.Get("upyun.host")
	user = config.Get("upyun.user")
	password = config.Get("upyun.password")
}


func UpyunUpload(read io.ReadCloser,path string)string{
	uploadPath := nameSpace+path;
	client := &http.Client{}
	req, _ := http.NewRequest("PUT", host+uploadPath, read)
	req.SetBasicAuth(user, password)
	resp,err := client.Do(req)
	if err!=nil {
		log.Error("upyun upload error",zap.Error(err))
		return ""
	}
	respBody,_ := ioutil.ReadAll(resp.Body)
	log.Info("upyun upload",zap.String("resp",string(respBody)))
	return downloadHost+"/"+path
}
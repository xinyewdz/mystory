package util

import (
	"bufio"
	"io"
	"net/http"
	"os"
	"strings"
)

var(
	downloadHost = "http://up.wenqiuqiu.com"
	nameSpace = "aaronimage/"
	host = ""
	user = ""
	password = ""
)

func init(){
	f,_ := os.Open("conf/story.conf")
	reader := bufio.NewReader(f)
	confMap := make(map[string]string)
	for{
		line,_,err := reader.ReadLine()
		if err==io.EOF {
			break
		}
		linStr := string(line)
		datas := strings.Split(linStr,"=")
		confMap[datas[0]] = datas[1]
	}
	host = confMap["upyun.host"]
	user = confMap["upyun.user"]
	password = confMap["upyun.password"]
}

func UpyunUpload(read io.ReadCloser,path string)string{
	uploadPath := nameSpace+path;
	client := &http.Client{}
	req, _ := http.NewRequest("POST", host+uploadPath, read)
	req.SetBasicAuth(user, password)
	client.Do(req)
	return downloadHost+path
}
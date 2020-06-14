package web

import (
	"context"
	"encoding/json"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"story-api/common"
	"story-api/store/dao/mongodao"
	"story-api/store/entity"
	"story-api/util"
	"strings"
)

var(
	storyDao = mongodao.NewStoryDao()
)

type StoryWeb struct {
}

func (web *StoryWeb)Upload(context context.Context,resp http.ResponseWriter,req *http.Request)*common.ApiResponse{
	req.ParseMultipartForm(1024*1024*100)
	file,fh,_ := req.FormFile("file")
	fileName := fh.Filename
	name := req.Form.Get("name")
	fileType := fileName[strings.LastIndex(fileName,".")+1:]
	uploadPath := "mystory/"+name+"."+fileType
	fUrl := util.UpyunUpload(file,uploadPath)
	ar := &common.ApiResponse{
		Data:"",
	}
	if fUrl==""{
		ar = common.Error("500","error")
	}else{
		ar = common.Success(fUrl)
	}
	return ar
}

func (web *StoryWeb)Save(context context.Context,resp http.ResponseWriter,req *http.Request)*common.ApiResponse{
	data,_ := ioutil.ReadAll(req.Body)
	mainLog.Info("savestory",zap.String("model",string(data)))
	storyObj := &entity.DBStory{}
	reqMap := make(map[string]string)
	json.Unmarshal(data,&reqMap)
	storyObj.Name = reqMap["name"]
	storyObj.AudioUrl = reqMap["audio"]
	storyObj.ImageUrl = reqMap["image"]
	storyDao.Insert(storyObj)
	return common.Success(storyObj.Id)
}



func (web *StoryWeb)List(context context.Context,resp http.ResponseWriter,req *http.Request)*common.ApiResponse{
	sl := storyDao.List()
	return common.Success(sl)
}

func (web *StoryWeb)Remove(context context.Context,resp http.ResponseWriter,req *http.Request)*common.ApiResponse{
	req.ParseForm()
	data,_ := ioutil.ReadAll(req.Body)
	reqMap := make(map[string]string)
	json.Unmarshal(data,&reqMap)
	storyDao.Remove(reqMap["id"])
	return common.Success(nil)
}

func (web *StoryWeb)Detail(context context.Context,resp http.ResponseWriter,req *http.Request)*common.ApiResponse{
	data,_ := ioutil.ReadAll(req.Body)
	reqMap := make(map[string]string)
	json.Unmarshal(data,&reqMap)
	idStr := reqMap["id"]
	sObj := storyDao.Get(idStr)
	if sObj==nil{
		return common.Error("40004","empty")
	}
	result := make(map[string]string)
	result["name"] = sObj.Name
	result["url"] = sObj.AudioUrl
	result["image"] = sObj.ImageUrl
	result["id"] = sObj.Id
	return common.Success(result)
}

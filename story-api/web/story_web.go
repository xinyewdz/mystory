package web

import (
	"context"
	"encoding/json"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"story-api/common"
	"story-api/store/dao/mysql"
	"story-api/store/entity"
	"story-api/util"
	"strings"
)

var(
	storyDao = mysql.NewStoryDao()
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
		ar.Error("500","error")
	}else{
		ar.Success(fUrl)
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
	ar := &common.ApiResponse{
		Data:0,
	}
	ar.Success(storyObj.Id)
	return ar
}



func (web *StoryWeb)List(context context.Context,resp http.ResponseWriter,req *http.Request)*common.ApiResponse{
	ar := &common.ApiResponse{}
	sl := storyDao.List()
	ar.Success(sl)
	return ar
}

func (web *StoryWeb)Remove(context context.Context,resp http.ResponseWriter,req *http.Request)*common.ApiResponse{
	req.ParseForm()
	data,_ := ioutil.ReadAll(req.Body)
	reqMap := make(map[string]string)
	json.Unmarshal(data,&reqMap)
	storyDao.Remove(reqMap["id"])
	ar := &common.ApiResponse{
	}
	ar.Success(nil)
	return ar
}

func (web *StoryWeb)Detail(context context.Context,resp http.ResponseWriter,req *http.Request)*common.ApiResponse{
	data,_ := ioutil.ReadAll(req.Body)
	reqMap := make(map[string]string)
	json.Unmarshal(data,&reqMap)
	idStr := reqMap["id"]
	sObj := storyDao.Detail(idStr)
	ar := &common.ApiResponse{}
	result := make(map[string]string)
	result["name"] = sObj.Name
	result["url"] = sObj.AudioUrl
	result["image"] = sObj.ImageUrl
	result["id"] = sObj.Id
	ar.Success(result)
	return ar
}

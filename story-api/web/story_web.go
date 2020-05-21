package web

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"story-api/common"
	"story-api/store/dao/leveldb"
	"story-api/store/entity"
	"story-api/util"
	"strconv"
	"strings"
)

var(
	storyDao = new(leveldb.StoryDao)
)

type StoryWeb struct {
}

func (web *StoryWeb)Upload(resp http.ResponseWriter,req *http.Request){
	req.ParseMultipartForm(1024*1024*100)
	file,fh,_ := req.FormFile("file")
	fileName := fh.Filename
	name := req.Form.Get("name")
	fileType := fileName[strings.LastIndex(fileName,".")+1:]
	uploadPath := "/mystory/"+name+"."+fileType
	fUrl := util.UpyunUpload(file,uploadPath)
	ar := &common.ApiResponse{
		Data:"",
	}
	if fUrl==""{
		ar.Error("500","error")
	}else{
		ar.Success(fUrl)
	}
	data,_ := json.Marshal(ar)
	resp.Write(data)
}

func (web *StoryWeb)Save(resp http.ResponseWriter,req *http.Request){
	data,_ := ioutil.ReadAll(req.Body)
	storyObj := &entity.DBStory{}
	json.Unmarshal(data,storyObj)
	storyDao.Insert(storyObj)
	ar := &common.ApiResponse{
		Data:0,
	}
	ar.Success(storyObj.Id)
	respData,_ := json.Marshal(ar)
	resp.Write(respData)
}



func (web *StoryWeb)List(resp http.ResponseWriter,req *http.Request){
	ar := &common.ApiResponse{}
	sl := storyDao.List()
	ar.Success(sl)
	data,_ := json.Marshal(ar)
	resp.Write(data)
}

func (web *StoryWeb)Remove(resp http.ResponseWriter,req *http.Request){
	req.ParseForm()
	idStr := req.Form.Get("id")
	id,_ := strconv.Atoi(idStr)
	storyDao.Remove(int64(id))
	ar := &common.ApiResponse{
	}
	ar.Success(nil)
	data,_ := json.Marshal(ar)
	resp.Write(data)
}

func (web *StoryWeb)Detail(resp http.ResponseWriter,req *http.Request){
	req.ParseForm()
	idStr := req.Form.Get("id")
	id,_ := strconv.Atoi(idStr)
	sObj := storyDao.Detail(int64(id))
	ar := &common.ApiResponse{}
	ar.Success(sObj)
	data,_ := json.Marshal(ar)
	resp.Write(data)
}

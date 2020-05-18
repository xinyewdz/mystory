package web

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"story-api/common"
	"story-api/store/dao/leveldb"
	"story-api/store/entity"
	"strconv"
)

var(
	userDao = new(leveldb.UserDao)
)

type UserWeb struct {
}

func (web *UserWeb)login(resp http.ResponseWriter,req *http.Request){

}

func (web *UserWeb)Save(resp http.ResponseWriter,req *http.Request){
	data,_ := ioutil.ReadAll(req.Body)
	storyObj := &entity.DBUser{}
	json.Unmarshal(data,storyObj)
	userDao.Insert(storyObj)
	ar := &common.ApiResponse{
		Data:0,
	}
	ar.Success(storyObj.Id)
	respData,_ := json.Marshal(ar)
	resp.Write(respData)
}



func (web *UserWeb)List(resp http.ResponseWriter,req *http.Request){
	ar := &common.ApiResponse{}
	sl := userDao.List()
	ar.Success(sl)
	data,_ := json.Marshal(ar)
	resp.Write(data)
}

func (web *UserWeb)Remove(resp http.ResponseWriter,req *http.Request){
	req.ParseForm()
	idStr := req.Form.Get("id")
	id,_ := strconv.Atoi(idStr)
	userDao.Remove(int64(id))
	ar := &common.ApiResponse{
	}
	ar.Success(nil)
	data,_ := json.Marshal(ar)
	resp.Write(data)
}

func (web *UserWeb)Detail(resp http.ResponseWriter,req *http.Request){
	req.ParseForm()
	idStr := req.Form.Get("id")
	id,_ := strconv.Atoi(idStr)
	sObj := userDao.Detail(int64(id))
	ar := &common.ApiResponse{}
	ar.Success(sObj)
	data,_ := json.Marshal(ar)
	resp.Write(data)
}

func  (web *UserWeb)AdminName(resp http.ResponseWriter,req *http.Request){
	ar := &common.ApiResponse{
	}
	ar.Success("aaron");
	data,_ := json.Marshal(ar)
	resp.Write(data)
}

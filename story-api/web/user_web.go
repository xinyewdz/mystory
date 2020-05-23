package web

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"story-api/common"
	"story-api/store/dao/leveldb"
	"story-api/store/entity"
	"story-api/util"
	"story-api/web/model"
	"strconv"
	"time"
)

var(
	userDao = new(leveldb.UserDao)
)

type UserWeb struct {
}

func (web *UserWeb)Login(context context.Context,resp http.ResponseWriter,req *http.Request)*common.ApiResponse{
	reqBody := make(map[string]string)
	resolveBody(req,&reqBody)
	code := reqBody["code"]
	sResp := util.Code2Session(code)
	apiResp := &common.ApiResponse{}
	if sResp==nil{
		apiResp.Error("501","getSession error")
		return apiResp
	}
	user := userDao.GetByOpenId(sResp.OpenId)
	if user==nil{
		user = &entity.DBUser{}
		user.Openid = sResp.OpenId
	}
	user.Name = reqBody["nickName"]
	user.Gender = reqBody["gender"]
	user.Phone = reqBody["phone"]

	if user.Id==0{
		userDao.Insert(user)
	}else{
		userDao.Update(user)
	}
	accountRes := &model.AccountResp{}
	util.CopyProperties(user,accountRes)
	token := strconv.Itoa(time.Now().Nanosecond())
	tokenMap[token]=user
	accountRes.Token = token
	apiResp.Success(accountRes)
	return apiResp
}

func (web *UserWeb)UpdateAdmin(context context.Context,resp http.ResponseWriter,req *http.Request)*common.ApiResponse{
	reqBody := make(map[string]string)
	resolveBody(req,&reqBody)
	userIdStr := reqBody["userId"]
	userType := reqBody["type"];
	userId,_ := strconv.Atoi(userIdStr)
	user := userDao.Get(int64(userId))
	if user!=nil{
		user.Type=userType
		userDao.Update(user)
	}
	return new(common.ApiResponse).Success(nil)
}

func (web *UserWeb)Save(context context.Context,resp http.ResponseWriter,req *http.Request)*common.ApiResponse{
	data,_ := ioutil.ReadAll(req.Body)
	storyObj := &entity.DBUser{}
	json.Unmarshal(data,storyObj)
	userDao.Insert(storyObj)
	ar := &common.ApiResponse{
		Data:0,
	}
	ar.Success(storyObj.Id)
	return ar
}



func (web *UserWeb)List(context context.Context,resp http.ResponseWriter,req *http.Request)*common.ApiResponse{
	ar := &common.ApiResponse{}
	sl := userDao.List()
	ar.Success(sl)
	return ar
}

func (web *UserWeb)Remove(context context.Context,resp http.ResponseWriter,req *http.Request)*common.ApiResponse{
	req.ParseForm()
	idStr := req.Form.Get("id")
	id,_ := strconv.Atoi(idStr)
	userDao.Remove(int64(id))
	ar := &common.ApiResponse{
	}
	ar.Success(nil)
	return ar
}

func (web *UserWeb)Detail(context context.Context,resp http.ResponseWriter,req *http.Request)*common.ApiResponse{
	req.ParseForm()
	idStr := req.Form.Get("id")
	id,_ := strconv.Atoi(idStr)
	sObj := userDao.Get(int64(id))
	ar := &common.ApiResponse{}
	ar.Success(sObj)
	return ar
}


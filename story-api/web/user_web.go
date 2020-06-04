package web

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"story-api/common"
	"story-api/store/dao/mongodao"
	"story-api/store/entity"
	"story-api/util"
	"story-api/web/model"
	"strconv"
	"time"
)

var(
	userDao = mongodao.NewUserDao()
)

type UserWeb struct {
}

func (web *UserWeb)Login(c context.Context,resp http.ResponseWriter,req *http.Request)*common.ApiResponse{
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

	if user.Id==""{
		user.Type = "user"
		userDao.Insert(user)
	}else{
		userDao.Update(user)
	}
	accountRes := &model.AccountResp{}
	util.CopyProperties(user,accountRes)
	token := strconv.Itoa(int(time.Now().Unix()))
	key := TOKEN_KEY+token
	ctx,_ := context.WithTimeout(context.Background(),5*time.Second)
	valData,_ := json.Marshal(user)
	res := redisClient.Set(ctx,key,valData,30*time.Minute)
	mainLog.Info(res.Val())
	accountRes.Token = token
	apiResp.Success(accountRes)
	return apiResp
}

func (web *UserWeb)Save(c context.Context,resp http.ResponseWriter,req *http.Request)*common.ApiResponse{
	ar := &common.ApiResponse{}
	user := c.Value(USER_KEY).(*entity.DBUser)
	if !isAdmin(user){
		ar.Error("100010","无权限")
		return ar
	}
	data,_ := ioutil.ReadAll(req.Body)
	storyObj := &entity.DBUser{}
	json.Unmarshal(data,storyObj)
	userDao.Insert(storyObj)
	ar.Success(storyObj.Id)
	return ar
}



func (web *UserWeb)List(c context.Context,resp http.ResponseWriter,req *http.Request)*common.ApiResponse{
	ar := &common.ApiResponse{}
	user := c.Value(USER_KEY).(*entity.DBUser)
	if !isAdmin(user){
		ar.Error("100010","无权限")
		return ar
	}
	sl := userDao.List()
	ar.Success(sl)
	return ar
}

func (web *UserWeb)Remove(c context.Context,resp http.ResponseWriter,req *http.Request)*common.ApiResponse{
	ar := &common.ApiResponse{}
	user := c.Value(USER_KEY).(*entity.DBUser)
	if !isAdmin(user){
		ar.Error("100010","无权限")
		return ar
	}
	req.ParseForm()
	idStr := req.Form.Get("id")
	userDao.Remove(idStr)
	ar.Success(nil)
	return ar
}

func (web *UserWeb)Detail(c context.Context,resp http.ResponseWriter,req *http.Request)*common.ApiResponse{
	ar := &common.ApiResponse{}
	user := c.Value(USER_KEY).(*entity.DBUser)
	if !isAdmin(user){
		ar.Error("100010","无权限")
		return ar
	}
	req.ParseForm()
	idStr := req.Form.Get("id")
	sObj := userDao.Get(idStr)
	ar.Success(sObj)
	return ar
}

func isAdmin(user *entity.DBUser)bool{
	return user!=nil&&user.Type=="admin"
}
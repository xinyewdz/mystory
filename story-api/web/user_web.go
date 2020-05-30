package web

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"story-api/common"
	"story-api/store/dao/mongo"
	"story-api/store/entity"
	"story-api/util"
	"story-api/web/model"
	"strconv"
	"time"
)

var(
	userDao = mongo.NewUserDao()
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
		userDao.Insert(user)
	}else{
		userDao.Update(user)
	}
	accountRes := &model.AccountResp{}
	util.CopyProperties(user,accountRes)
	token := strconv.Itoa(time.Now().Nanosecond())
	key := TOKEN_KEY+token
	ctx,_ := context.WithTimeout(context.Background(),5*time.Second)
	valData,_ := json.Marshal(user)
	redisClient.Set(ctx,key,valData,30*time.Second)
	accountRes.Token = token
	apiResp.Success(accountRes)
	return apiResp
}

func (web *UserWeb)UpdateAdmin(c context.Context,resp http.ResponseWriter,req *http.Request)*common.ApiResponse{
	reqBody := make(map[string]string)
	resolveBody(req,&reqBody)
	userIdStr := reqBody["userId"]
	userType := reqBody["type"];
	user := userDao.Get(userIdStr)
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
	userDao.Remove(idStr)
	ar := &common.ApiResponse{
	}
	ar.Success(nil)
	return ar
}

func (web *UserWeb)Detail(context context.Context,resp http.ResponseWriter,req *http.Request)*common.ApiResponse{
	req.ParseForm()
	idStr := req.Form.Get("id")
	sObj := userDao.Get(idStr)
	ar := &common.ApiResponse{}
	ar.Success(sObj)
	return ar
}


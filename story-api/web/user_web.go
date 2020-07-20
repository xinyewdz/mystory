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

var (
	userDao = mongodao.NewUserDao()
)

type UserWeb struct {
	EnableStoryList bool
}

func (web *UserWeb) Login(c context.Context, resp http.ResponseWriter, req *http.Request) *common.ApiResponse {
	reqBody := make(map[string]string)
	resolveBody(req, &reqBody)
	code := reqBody["code"]
	sResp := util.Code2Session(code)
	if sResp == nil {
		return common.Error("501", "getSession error")
	}
	user := userDao.GetByOpenId(sResp.OpenId)
	if user == nil {
		user = &entity.DBUser{}
		user.Openid = sResp.OpenId
	}
	user.Name = reqBody["nickName"]
	user.Gender = reqBody["gender"]
	user.Phone = reqBody["phone"]
	user.AvatarUrl = reqBody["avatarUrl"]
	user.Province = reqBody["province"]
	user.City = reqBody["city"]

	if user.Id == "" {
		user.Type = "user"
		userDao.Insert(user)
	} else {
		userDao.Update(user)
	}
	accountRes := &model.AccountResp{}
	accountRes.EnableStoryList = ((user.Type == entity.USER_TYPE_ADMIN) || web.EnableStoryList)
	util.CopyProperties(user, accountRes)
	token := strconv.Itoa(int(time.Now().Unix()))
	key := TOKEN_KEY + token
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	valData, _ := json.Marshal(user)
	res := redisClient.Set(ctx, key, valData, 240*time.Hour)
	mainLog.Info(res.Val())
	accountRes.Token = token
	return common.Success(accountRes)
}

func (web *UserWeb) Save(c context.Context, resp http.ResponseWriter, req *http.Request) *common.ApiResponse {
	user := c.Value(USER_KEY).(*entity.DBUser)
	if !isAdmin(user) {
		return common.Error("100010", "无权限")
	}
	data, _ := ioutil.ReadAll(req.Body)
	userObj := &entity.DBUser{}
	json.Unmarshal(data, userObj)
	if userObj.Id == "" {
		userDao.Insert(userObj)
	} else {
		userDao.Update(userObj)
	}
	return common.Success(userObj.Id)
}

func (web *UserWeb) List(c context.Context, resp http.ResponseWriter, req *http.Request) *common.ApiResponse {
	user := c.Value(USER_KEY).(*entity.DBUser)
	if !isAdmin(user) {
		return common.Error("100010", "无权限")
	}
	sl := userDao.List()
	return common.Success(sl)
}

func (web *UserWeb) Remove(c context.Context, resp http.ResponseWriter, req *http.Request) *common.ApiResponse {
	user := c.Value(USER_KEY).(*entity.DBUser)
	if !isAdmin(user) {
		return common.Error("100010", "无权限")
	}
	reqMap := parseBody(req)
	idStr := reqMap["id"]
	userDao.Remove(idStr)
	return common.Success(nil)
}

func (web *UserWeb) Detail(c context.Context, resp http.ResponseWriter, req *http.Request) *common.ApiResponse {
	user := c.Value(USER_KEY).(*entity.DBUser)
	if !isAdmin(user) {
		return common.Error("100010", "无权限")
	}
	reqMap := parseBody(req)
	sObj := userDao.Get(reqMap["id"])
	return common.Success(sObj)
}

func isAdmin(user *entity.DBUser) bool {
	return user != nil && user.Type == "admin"
}

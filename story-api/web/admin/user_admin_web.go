package admin

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"story-api/common"
	"story-api/store/entity"
	"story-api/util"
	"story-api/web"
	"story-api/web/model"
	"strconv"
	"time"
)

type UserAdminWeb struct {
	EnableStoryList bool
}

func (userweb *UserAdminWeb) Login(c context.Context, resp http.ResponseWriter, req *http.Request) *common.ApiResponse {
	reqBody := make(map[string]string)
	web.ResolveBody(req, &reqBody)
	phone := reqBody["phone"]
	passwd := reqBody["password"]
	user := userDao.GetByPhone(phone)
	if user == nil || user.Password != passwd {
		return common.Error("4002", "用户名或密码错误")
	}
	if user.Type != entity.USER_TYPE_ADMIN {
		return common.Error("4003", "无权限登录")
	}
	accountRes := &model.AccountResp{}
	util.CopyProperties(user, accountRes)
	token := strconv.Itoa(int(time.Now().Unix()))
	key := web.TOKEN_KEY + token
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	valData, _ := json.Marshal(user)
	res := redisClient.Set(ctx, key, valData, 240*time.Hour)
	mainLog.Info(res.Val())
	accountRes.Token = token
	return common.Success(accountRes)
}

func (web *UserAdminWeb) ListUserPlayDetail(c context.Context, resp http.ResponseWriter, req *http.Request) *common.ApiResponse {
	data, _ := ioutil.ReadAll(req.Body)
	reqMap := make(map[string]string)
	json.Unmarshal(data, &reqMap)
	idStr := reqMap["id"]
	list := storyPlayDetailDao.List("", idStr)
	smap := make(map[string]*entity.DBStory)
	respList := []*model.StoryPlayDetailResp{}
	for _, obj := range list {
		s := &model.StoryPlayDetailResp{}
		s.PlayTime = obj.CreateTime.Format("2006-01-02 15:04:05")
		story, ok := smap[obj.StoryId]
		if !ok {
			story = storyDao.Get(obj.StoryId)
			if story != nil {
				smap[obj.StoryId] = story
			}
		}
		if story != nil {
			s.Id = story.Id
			s.Name = story.Name
		} else {
			s.Name = "-"
			s.Id = ""
		}
		respList = append(respList, s)
	}
	return common.Success(respList)

}

func (web *UserAdminWeb) ListStoryPlayDetail(c context.Context, resp http.ResponseWriter, req *http.Request) *common.ApiResponse {
	data, _ := ioutil.ReadAll(req.Body)
	reqMap := make(map[string]string)
	json.Unmarshal(data, &reqMap)
	idStr := reqMap["id"]
	list := storyPlayDetailDao.List(idStr, "")
	accountMap := make(map[string]*entity.DBUser)
	respList := []*model.StoryPlayDetailResp{}
	for _, obj := range list {
		s := &model.StoryPlayDetailResp{}
		s.PlayTime = obj.CreateTime.Format("2006-01-02 15:04:05")
		a, ok := accountMap[obj.UserId]
		if !ok {
			a = userDao.Get(obj.UserId)
			if a != nil {
				accountMap[obj.UserId] = a
			}
		}
		if a != nil {
			s.Id = a.Id
			s.Name = a.Name
		} else {
			s.Name = "-"
			s.Id = ""
		}
		respList = append(respList, s)
	}
	return common.Success(respList)

}

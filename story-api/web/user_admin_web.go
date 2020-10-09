package web

import (
	"context"
	"encoding/json"
	"net/http"
	"story-api/common"
	"story-api/store/entity"
	"story-api/util"
	"story-api/web/model"
	"strconv"
	"time"
)

type UserAdminWeb struct {
	EnableStoryList bool
}

func (web *UserAdminWeb) Login(c context.Context, resp http.ResponseWriter, req *http.Request) *common.ApiResponse {
	reqBody := make(map[string]string)
	resolveBody(req, &reqBody)
	phone := reqBody["phone"]
	passwd := reqBody["password"]
	user := userDao.GetByPhone(phone)
	if user == nil || user.Password != passwd {
		return nil
	}
	accountRes := &model.AccountResp{}
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

func (web *UserAdminWeb) ListStoryPlayDetail(c context.Context, resp http.ResponseWriter, req *http.Request) *common.ApiResponse {
	user := c.Value(USER_KEY).(*entity.DBUser)
	list := storyPlayDetailDao.List("", user.Id)
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

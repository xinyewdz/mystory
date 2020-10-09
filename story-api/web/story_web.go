package web

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"story-api/common"
	"story-api/store/entity"
	"story-api/util"
	"story-api/web/model"
	"strings"

	"go.uber.org/zap"
)

type StoryWeb struct {
}

func (web *StoryWeb) Upload(context context.Context, resp http.ResponseWriter, req *http.Request) *common.ApiResponse {
	req.ParseMultipartForm(1024 * 1024 * 100)
	file, fh, _ := req.FormFile("file")
	fileName := fh.Filename
	name := req.Form.Get("name")
	fileType := fileName[strings.LastIndex(fileName, ".")+1:]
	uploadPath := "mystory/" + name + "." + fileType
	fUrl := util.UpyunUpload(file, uploadPath)
	ar := &common.ApiResponse{
		Data: "",
	}
	if fUrl == "" {
		ar = common.Error("500", "error")
	} else {
		ar = common.Success(fUrl)
	}
	return ar
}

func (web *StoryWeb) Play(c context.Context, resp http.ResponseWriter, req *http.Request) *common.ApiResponse {
	user, ok := c.Value(USER_KEY).(*entity.DBUser)
	var userId string
	if !ok || user == nil {
		userId = "0"
	} else {
		userId = user.Id
	}
	data, _ := ioutil.ReadAll(req.Body)
	reqMap := make(map[string]string)
	json.Unmarshal(data, &reqMap)
	storyId := reqMap["storyId"]
	detail := &entity.DBStoryPlayDetail{
		StoryId: storyId,
		UserId:  userId,
	}
	storyPlayDetailDao.Insert(detail)
	return common.Success(nil)
}

func (web *StoryWeb) Save(context context.Context, resp http.ResponseWriter, req *http.Request) *common.ApiResponse {
	data, _ := ioutil.ReadAll(req.Body)
	mainLog.Info("savestory", zap.String("model", string(data)))
	storyObj := &entity.DBStory{}
	json.Unmarshal(data, storyObj)
	if storyObj.Id != "" {
		storyDao.UpdateObj(storyObj.Id, storyObj)
	} else {
		storyDao.Insert(storyObj)
	}
	return common.Success(storyObj.Id)
}

func (web *StoryWeb) PlayList(c context.Context, resp http.ResponseWriter, req *http.Request) *common.ApiResponse {
	cuser := c.Value(USER_KEY)
	var user *entity.DBUser
	if cuser != nil {
		user = cuser.(*entity.DBUser)
	}
	var sl []*entity.DBStory

	flag := true
	sl = storyDao.List(&flag, "")
	if user != nil {
		slMy := storyDao.List(nil, user.Id)
		for _, s := range slMy {
			exist := isExist(s, sl)
			if !exist {
				sl = append(sl, s)
			}
		}
	}
	list := []*model.StoryPlayResp{}
	for _, s := range sl {
		resp := &model.StoryPlayResp{
			Id:       s.Id,
			Name:     s.Name,
			ImageUrl: s.ImageUrl,
			AudioUrl: s.AudioUrl,
		}
		user := userDao.Get(s.CreateUser)
		if user != nil {
			resp.CreateUser = user.Name
		} else {
			resp.CreateUser = "-"
		}
		resp.CreateTime = s.CreateTime.Format("2006-01-02")
		total := storyPlayDetailDao.Count(s.Id, "")
		resp.TotalPlay = total
		list = append(list, resp)

	}

	return common.Success(list)
}

func (web *StoryWeb) List(c context.Context, resp http.ResponseWriter, req *http.Request) *common.ApiResponse {
	user, ok := c.Value(USER_KEY).(*entity.DBUser)
	var sl []*entity.DBStory
	if ok && user.Type == entity.USER_TYPE_ADMIN {
		sl = storyDao.List(nil, "")
	} else {
		sl = storyDao.List(nil, user.Id)
	}
	return common.Success(sl)
}

func (web *StoryWeb) Remove(context context.Context, resp http.ResponseWriter, req *http.Request) *common.ApiResponse {
	req.ParseForm()
	data, _ := ioutil.ReadAll(req.Body)
	reqMap := make(map[string]string)
	json.Unmarshal(data, &reqMap)
	storyDao.Remove(reqMap["id"])
	return common.Success(nil)
}

func (web *StoryWeb) Detail(context context.Context, resp http.ResponseWriter, req *http.Request) *common.ApiResponse {
	data, _ := ioutil.ReadAll(req.Body)
	reqMap := make(map[string]string)
	json.Unmarshal(data, &reqMap)
	idStr := reqMap["id"]
	sObj := storyDao.Get(idStr)
	if sObj == nil {
		return common.Error("40004", "empty")
	}
	return common.Success(sObj)
}

func (web *StoryWeb) AddFavorite(c context.Context, resp http.ResponseWriter, req *http.Request) *common.ApiResponse {
	user, _ := c.Value(USER_KEY).(*entity.DBUser)
	userId := user.Id
	data, _ := ioutil.ReadAll(req.Body)
	reqMap := make(map[string]string)
	json.Unmarshal(data, &reqMap)
	storyId := reqMap["storyId"]
	favorite := &entity.DBStoryFavorite{
		StoryId: storyId,
		UserId:  userId,
	}
	storyFavoriteDao.Insert(favorite)
	return common.Success("")
}

func (web *StoryWeb) RemoveFavorite(c context.Context, resp http.ResponseWriter, req *http.Request) *common.ApiResponse {
	user, _ := c.Value(USER_KEY).(*entity.DBUser)
	userId := user.Id
	data, _ := ioutil.ReadAll(req.Body)
	reqMap := make(map[string]string)
	json.Unmarshal(data, &reqMap)
	storyId := reqMap["storyId"]
	storyFavoriteDao.Remove(userId, storyId)
	return common.Success("")
}

func (web *StoryWeb) ListFavorite(c context.Context, resp http.ResponseWriter, req *http.Request) *common.ApiResponse {
	user, _ := c.Value(USER_KEY).(*entity.DBUser)
	userId := user.Id
	fl := storyFavoriteDao.List(userId, "")
	sl := make([]*entity.DBStory, len(fl))
	for _, f := range fl {
		s := storyDao.Get(f.StoryId)
		if s != nil {
			sl = append(sl, s)
		}
	}
	return common.Success(sl)
}

func isExist(s *entity.DBStory, sl []*entity.DBStory) bool {
	for _, _s := range sl {
		if _s.Id == s.Id {
			return true
		}
	}
	return false
}

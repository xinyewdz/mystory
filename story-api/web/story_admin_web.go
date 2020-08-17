package web

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"story-api/common"
	"story-api/store/entity"
	"story-api/web/model"
)

type StoryAdminWeb struct {
}

func (web *StoryAdminWeb) ListUserPlayDetail(c context.Context, resp http.ResponseWriter, req *http.Request) *common.ApiResponse {
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

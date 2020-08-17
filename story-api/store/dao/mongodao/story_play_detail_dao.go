package mongodao

import (
	"story-api/store/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type StoryPlayDetailDao struct {
	BaseDao
}

func NewStoryPlayDetailDao() *StoryPlayDetailDao {
	dao := &StoryPlayDetailDao{}
	dao.Table = "story_play_detail"
	dao.Obj = entity.DBStoryPlayDetail{}
	return dao
}

func (dao *StoryPlayDetailDao) Insert(obj *entity.DBStoryPlayDetail) {
	obj.CreateTime = time.Now()
	obj.Id = primitive.NewObjectID().Hex()
	dao.InsertObj(obj)
}

func (dao *StoryPlayDetailDao) List(storyId string, userId string) []*entity.DBStoryPlayDetail {
	filter := make(map[string]interface{})
	if storyId != "" {
		filter["storyId"] = storyId
	}
	if userId != "" {
		filter["userId"] = userId
	}
	orderOpt := &options.FindOptions{
		Sort: bson.M{
			"createTime": -1,
		},
	}
	objs := dao.ListByFilter(filter, orderOpt)
	if objs == nil {
		return nil
	}
	list := []*entity.DBStoryPlayDetail{}
	for _, obj := range objs {
		list = append(list, obj.(*entity.DBStoryPlayDetail))
	}
	return list
}

func (dao *StoryPlayDetailDao) Count(storyId string, userId string) int64 {
	filter := make(map[string]interface{})
	if storyId != "" {
		filter["storyId"] = storyId
	}
	if userId != "" {
		filter["userId"] = userId
	}
	total := dao.CountByFilter(filter)

	return total
}

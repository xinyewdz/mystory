package mongodao

import (
	"story-api/store/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StoryFavoriteDao struct {
	BaseDao
}

func NewStoryFavoriteDao() *StoryFavoriteDao {
	dao := &StoryFavoriteDao{}
	dao.Table = "story_favorite"
	dao.Obj = entity.DBStoryFavorite{}
	return dao
}

func (dao *StoryFavoriteDao) Insert(obj *entity.DBStoryFavorite) {
	obj.CreateTime = time.Now()
	obj.Id = primitive.NewObjectID().Hex()
	dao.InsertObj(obj)
}

func (dao *StoryFavoriteDao) List(userId string, storyId string) []*entity.DBStoryFavorite {
	filter := make(map[string]interface{})
	if userId != "" {
		filter["userId"] = userId
	}
	if storyId != "" {
		filter["storyId"] = storyId
	}
	objs := dao.ListByFilter(filter)
	if objs == nil {
		return nil
	}
	list := []*entity.DBStoryFavorite{}
	for _, obj := range objs {
		list = append(list, obj.(*entity.DBStoryFavorite))
	}
	return list
}

func (dao *StoryFavoriteDao) Remove(userId string, storyId string) int64 {
	filter := make(map[string]interface{})
	if userId != "" {
		filter["userId"] = userId
	}
	if storyId != "" {
		filter["storyId"] = storyId
	}
	count := dao.RemoveObjByFilter(filter)
	return count
}

package mongodao

import (
	"story-api/store/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type StoryDao struct {
	BaseDao
}

func NewStoryDao() *StoryDao {
	dao := &StoryDao{}
	dao.Table = "story"
	dao.Obj = entity.DBStory{}
	return dao
}

func (dao *StoryDao) Insert(obj *entity.DBStory) {
	obj.CreateTime = time.Now()
	obj.Id = primitive.NewObjectID().Hex()
	dao.InsertObj(obj)
}

func (dao *StoryDao) Update(obj *entity.DBStory) {
	id := obj.Id
	dao.UpdateObj(id, obj)
}

func (dao *StoryDao) List(isPublic *bool, createUser string) []*entity.DBStory {
	filter := make(map[string]interface{})
	if isPublic != nil {
		filter["isPublic"] = isPublic
	}

	if createUser != "" {
		filter["createUser"] = createUser
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
	list := []*entity.DBStory{}
	for _, obj := range objs {
		list = append(list, obj.(*entity.DBStory))
	}
	return list
}

func (dao *StoryDao) Get(id string) *entity.DBStory {
	obj := dao.GetObj(id)
	if obj != nil {
		return obj.(*entity.DBStory)
	}
	return nil
}

func (dao *StoryDao) Remove(id string) {
	dao.RemoveObj(id)
}

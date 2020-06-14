package mongodao

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"story-api/store/entity"
	"time"
)


type StoryDao struct {
	BaseDao
}

func NewStoryDao()*StoryDao{
	dao := &StoryDao{
	}
	dao.Table = "story"
	dao.Obj = entity.DBStory{}
	return dao
}

func (dao *StoryDao) Insert(obj *entity.DBStory){
	obj.CreateTime = time.Now()
	obj.Id = primitive.NewObjectID().Hex()
	dao.InsertObj(obj)
}

func (dao *StoryDao) Update(obj *entity.DBStory){
	id := obj.Id
	dao.UpdateObj(id,obj)
}

func (dao *StoryDao)  List()[]*entity.DBStory{
	objs := dao.ListAll()
	if objs==nil{
		return nil
	}
	list := []*entity.DBStory{}
	for _,obj:= range objs{
		list = append(list,obj.(*entity.DBStory))
	}
	return list
}

func (dao *StoryDao) Get(id string)*entity.DBStory{
	obj := dao.GetObj(id)
	if obj!=nil{
		return obj.(*entity.DBStory)
	}
	return nil
}

func (dao *StoryDao) Remove(id string){
	dao.RemoveObj(id)
}

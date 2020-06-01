package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"story-api/store/entity"
	"time"
)


type StoryDao struct {
	table string
}

func NewStoryDao()*StoryDao{
	dao := &StoryDao{
		table: "story",
	}
	return dao
}

func (dao *StoryDao) Insert(obj *entity.DBStory){
	obj.CreateTime = time.Now()
	obj.Id = primitive.NewObjectID().Hex()
	ctx,_ := context.WithTimeout(context.Background(),5*time.Second)
	doc := toDoc(obj)
	_,err := dClient.Collection(dao.table).InsertOne(ctx,doc)
	if err!=nil{
		panic(err)
	}
}

func (dao *StoryDao) Update(obj *entity.DBStory){
	ctx,_ := context.WithTimeout(context.Background(),5*time.Second)
	query := bson.M{
		"_id":obj.Id,
	}
	doc := toDoc(obj)
	update := bson.M{
		"$set":doc,
	}
	_,err := dClient.Collection(dao.table).UpdateOne(ctx,query,update)
	if err!=nil{
		panic(err)
	}
}

func (dao *StoryDao)  List()[]*entity.DBStory{
	sl := []*entity.DBStory{}
	ctx,_ := context.WithTimeout(context.Background(),5*time.Second)
	query := bson.M{
	}
	cursor,err := dClient.Collection(dao.table).Find(ctx,query)
	if err!=nil{
		log.Error("story query error",zap.Error(err))
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx){
		obj := &entity.DBStory{}
		cursor.Decode(obj)
		sl = append(sl,obj)
	}
	return sl
}

func (dao *StoryDao) Detail(id string)*entity.DBStory{
	ctx,_ := context.WithTimeout(context.Background(),5*time.Second)
	query := bson.M{
		"_id":id,
	}
	s := &entity.DBStory{}
	dClient.Collection(dao.table).FindOne(ctx,query).Decode(s)
	return s
}

func (dao *StoryDao) Remove(id string){
	ctx,_ := context.WithTimeout(context.Background(),5*time.Second)
	query := bson.M{
		"_id":id,
	}
	_,err := dClient.Collection(dao.table).DeleteOne(ctx,query)
	if err!=nil {
		panic(err)
	}
}

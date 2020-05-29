package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"story-api/store/entity"
	"time"
)


var storyDb *mongo.Database

type StoryDao struct {
	table string
}

func init(){
	url := "mongodb://"+host+"/"+database
	ctx,_ := context.WithTimeout(context.Background(),5*time.Second)
	client,err := mongo.Connect(ctx,options.Client().ApplyURI(url))
	if err!=nil{
		log.Error("connect mongo error.",zap.String("host",host),zap.Error(err))
	}
	storyDb = client.Database(database)
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
	_,err := storyDb.Collection(dao.table).InsertOne(ctx,doc)
	if err!=nil{
		panic(err)
	}
}

func (dao *StoryDao) Update(obj *entity.DBStory){
	ctx,_ := context.WithTimeout(context.Background(),5*time.Second)
	query := bson.M{
		"_id":obj.Id,
	}
	_,err := storyDb.Collection(dao.table).UpdateOne(ctx,query,obj)
	if err!=nil{
		panic(err)
	}
}

func (dao *StoryDao)  List()[]*entity.DBStory{
	sl := []*entity.DBStory{}
	ctx,_ := context.WithTimeout(context.Background(),5*time.Second)
	query := bson.M{
	}
	cursor,err := storyDb.Collection(dao.table).Find(ctx,query)
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
	storyDb.Collection(dao.table).FindOne(ctx,query).Decode(s)
	return s
}

func (dao *StoryDao) Remove(id string){
	ctx,_ := context.WithTimeout(context.Background(),5*time.Second)
	query := bson.M{
		"_id":id,
	}
	_,err := storyDb.Collection(dao.table).DeleteOne(ctx,query)
	if err!=nil {
		panic(err)
	}
}

package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"story-api/store/entity"
	"time"
)

type UserDao struct {
	table string
}

func NewUserDao()*UserDao{
	userdao := &UserDao{
		table: "user",
	}
	return userdao
}

func (dao *UserDao) Insert(obj *entity.DBUser){
	obj.CreateTime = time.Now()
	obj.Id = primitive.NewObjectID().Hex()
	ctx,_ := context.WithTimeout(context.Background(),5*time.Second)
	doc := toDoc(obj)
	dClient.Collection(dao.table).InsertOne(ctx,doc)
}

func (dao *UserDao) Update(obj *entity.DBUser){
	ctx,_ := context.WithTimeout(context.Background(),5*time.Second)
	query := bson.M{
		"_id":obj.Id,
	}
	_,err := dClient.Collection(dao.table).UpdateOne(ctx,query,obj)
	if err!=nil{
		panic(err)
	}
}

func (dao *UserDao) List()[]*entity.DBUser{
	sl := []*entity.DBUser{}
	ctx,_ := context.WithTimeout(context.Background(),5*time.Second)
	query := bson.M{
	}
	cursor,err := dClient.Collection(dao.table).Find(ctx,query)
	if err!=nil{
		panic(err)
		log.Error("user query error",zap.Error(err))
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx){
		obj := &entity.DBUser{}
		cursor.Decode(obj)
		sl = append(sl,obj)
	}
	return sl
}

func (dao *UserDao) Get(id string)*entity.DBUser{
	query := bson.M{
		"_id":id,
	}
	ctx,_ := context.WithTimeout(context.Background(),5*time.Second)
	result := dClient.Collection(dao.table).FindOne(ctx,query)
	if result.Err()!=nil{
		panic(result.Err())
		return nil
	}
	obj := &entity.DBUser{}
	result.Decode(obj)
	return obj
}

func (dao *UserDao) GetByOpenId(openId string)*entity.DBUser{
	query := bson.M{
		"openId":openId,
	}
	ctx,_ := context.WithTimeout(context.Background(),5*time.Second)
	result := dClient.Collection(dao.table).FindOne(ctx,query)
	if result.Err()!=nil{
		panic(result.Err())
		return nil
	}
	obj := &entity.DBUser{}
	result.Decode(obj)
	return obj
}

func (dao *UserDao) Remove(id string){
	query := bson.M{
		"_id":id,
	}
	ctx,_ := context.WithTimeout(context.Background(),5*time.Second)
	dClient.Collection(dao.table).DeleteOne(ctx,query)
}
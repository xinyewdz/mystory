package mongodao

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
	doc := toDoc(obj)
	update := bson.M{
		"$set":doc,
	}
	_,err := dClient.Collection(dao.table).UpdateOne(ctx,query,update)
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
	defer cursor.Close(ctx)
	if err==nil{
		for cursor.Next(ctx){
			obj := &entity.DBUser{}
			cursor.Decode(obj)
			sl = append(sl,obj)
		}
		return sl
	}
	if err==mongo.ErrNoDocuments{
		return nil
	}else{
		panic(err)
	}
}

func (dao *UserDao) Get(id string)*entity.DBUser{
	query := bson.M{
		"_id":id,
	}
	ctx,_ := context.WithTimeout(context.Background(),5*time.Second)
	result := dClient.Collection(dao.table).FindOne(ctx,query)
	obj := &entity.DBUser{}
	err := result.Decode(obj)
	if err==nil{
		return obj
	}
	if err==mongo.ErrNoDocuments{
		return nil
	}
	if err==mongo.ErrNoDocuments{
		return nil
	}else{
		panic(err)
	}
}

func (dao *UserDao) GetByName(name string)*entity.DBUser{
	query := bson.M{
		"name":name,
	}
	ctx,_ := context.WithTimeout(context.Background(),5*time.Second)
	result := dClient.Collection(dao.table).FindOne(ctx,query)
	obj := &entity.DBUser{}
	err := result.Decode(obj)
	if err==nil{
		return obj
	}
	if err==mongo.ErrNoDocuments{
		return nil
	}else{
		panic(err)
	}
}

func (dao *UserDao) GetByOpenId(openId string)*entity.DBUser{
	query := bson.M{
		"openId":openId,
	}
	ctx,_ := context.WithTimeout(context.Background(),5*time.Second)
	result := dClient.Collection(dao.table).FindOne(ctx,query)
	obj := &entity.DBUser{}
	err := result.Decode(obj)
	if err==nil{
		return obj
	}
	if err==mongo.ErrNoDocuments{
		return nil
	}else{
		panic(err)
	}
}

func (dao *UserDao) Remove(id string){
	query := bson.M{
		"_id":id,
	}
	ctx,_ := context.WithTimeout(context.Background(),5*time.Second)
	dClient.Collection(dao.table).DeleteOne(ctx,query)
}
package mongodao

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
	"time"
)


type BaseDao struct {
	Table string
	Obj interface{}
}

func (dao *BaseDao) InsertObj(obj interface{}){
	ctx,_ := context.WithTimeout(context.Background(),5*time.Second)
	doc := toDoc(obj)
	_,err := dClient.Collection(dao.Table).InsertOne(ctx,doc)
	if err!=nil{
		panic(err)
	}
}

func (dao *BaseDao) UpdateObj(id string,obj interface{}){
	ctx,_ := context.WithTimeout(context.Background(),5*time.Second)
	query := bson.M{
		"_id":id,
	}
	doc := toDoc(obj)
	update := bson.M{
		"$set":doc,
	}
	_,err := dClient.Collection(dao.Table).UpdateOne(ctx,query,update)
	if err!=nil{
		panic(err)
	}
}

func (dao *BaseDao)  ListAll()[]interface{}{
	t := reflect.TypeOf(dao.Obj)
	if t.Kind()==reflect.Ptr{
		t = t.Elem()
	}
	ctx,_ := context.WithTimeout(context.Background(),5*time.Second)
	query := bson.M{
	}
	list := []interface{}{}
	cursor,err := dClient.Collection(dao.Table).Find(ctx,query)
	if err==nil{
		defer cursor.Close(ctx)
		for cursor.Next(ctx){
			obj := reflect.New(t).Elem().Addr().Interface()
			cursor.Decode(obj)
			list = append(list,obj)
		}
		return list
	}
	if err==mongo.ErrNoDocuments{
		return nil
	}else{
		panic(err)
	}
}

func (dao *BaseDao) GetObj(id string)interface{}{
	ctx,_ := context.WithTimeout(context.Background(),5*time.Second)
	query := bson.M{
		"_id":id,
	}
	t := reflect.TypeOf(dao.Obj)
	if t.Kind()==reflect.Ptr{
		t = t.Elem()
	}
	obj := reflect.New(t).Elem().Addr().Interface()
	err := dClient.Collection(dao.Table).FindOne(ctx,query).Decode(&obj)
	if err==nil{
		return &obj
	}
	if err==mongo.ErrNoDocuments{
		return nil
	}else{
		panic(err)
	}
}

func (dao *BaseDao) RemoveObj(id string){
	ctx,_ := context.WithTimeout(context.Background(),5*time.Second)
	query := bson.M{
		"_id":id,
	}
	_,err := dClient.Collection(dao.Table).DeleteOne(ctx,query)
	if err!=nil {
		panic(err)
	}
}

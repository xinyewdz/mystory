package mongodao

import (
	"context"
	"reflect"
	"story-api/global"
	"story-api/util/config"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
	"go.uber.org/zap"
)

var (
	log     = global.MainLog
	dClient *mongo.Database
	tz, _   = time.LoadLocation("Local")
)

func init() {
	host := config.Get("mongodb.host")
	database := config.Get("mongodb.database")
	user := config.Get("mongodb.user")
	password := config.Get("mongodb.password")
	url := "mongodb://" + user + ":" + password + "@" + host + "/admin"
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		log.Error("connect mongodao error.", zap.String("host", host), zap.Error(err))
	}
	dClient = client.Database(database)
}

func toDoc(v interface{}) *bsoncore.Document {
	data, err := bson.Marshal(v)
	if err != nil {
		return nil
	}
	doc := &bsoncore.Document{}
	bson.Unmarshal(data, doc)
	return doc
}

func fromDoc(doc *bsoncore.Document, v interface{}) {
	data, err := bson.Marshal(doc)
	if err != nil {
		return
	}
	err = bson.Unmarshal(data, v)
	if err != nil {
		log.Error("fromDoc error", zap.Error(err))
	}
}

type BaseDao struct {
	Table string
	Obj   interface{}
}

func (dao *BaseDao) InsertObj(obj interface{}) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	doc := toDoc(obj)
	_, err := dClient.Collection(dao.Table).InsertOne(ctx, doc)
	if err != nil {
		panic(err)
	}
}

func (dao *BaseDao) UpdateObj(id string, obj interface{}) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	query := bson.M{
		"_id": id,
	}
	doc := toDoc(obj)
	update := bson.M{
		"$set": doc,
	}
	_, err := dClient.Collection(dao.Table).UpdateOne(ctx, query, update)
	if err != nil {
		panic(err)
	}
}

func (dao *BaseDao) ListAll() []interface{} {
	t := reflect.TypeOf(dao.Obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	query := bson.M{}
	list := []interface{}{}
	cursor, err := dClient.Collection(dao.Table).Find(ctx, query)
	if err == nil {
		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			obj := reflect.New(t).Elem().Addr().Interface()
			cursor.Decode(obj)
			list = append(list, obj)
		}
		return list
	}
	if err == mongo.ErrNoDocuments {
		return nil
	} else {
		panic(err)
	}
}

func (dao *BaseDao) ListByFilter(filter map[string]interface{}, opts ...*options.FindOptions) []interface{} {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	t := reflect.TypeOf(dao.Obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	query := bson.M{}
	for key, val := range filter {
		query[key] = val
	}
	list := []interface{}{}
	cursor, err := dClient.Collection(dao.Table).Find(ctx, query, opts...)
	if err == nil {
		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			obj := reflect.New(t).Elem().Addr().Interface()
			cursor.Decode(obj)
			list = append(list, obj)
		}
		return list
	}
	if err == mongo.ErrNoDocuments {
		return nil
	} else {
		panic(err)
	}
}

func (dao *BaseDao) GetObj(id string) interface{} {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	query := bson.M{
		"_id": id,
	}
	t := reflect.TypeOf(dao.Obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	obj := reflect.New(t).Elem().Addr().Interface()
	err := dClient.Collection(dao.Table).FindOne(ctx, query).Decode(obj)
	if err == nil {
		return obj
	}
	if err == mongo.ErrNoDocuments {
		return nil
	} else {
		panic(err)
	}
}

func (dao *BaseDao) RemoveObj(id string) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	query := bson.M{
		"_id": id,
	}
	_, err := dClient.Collection(dao.Table).DeleteOne(ctx, query)
	if err != nil {
		panic(err)
	}
}

func (dao *BaseDao) RemoveObjByFilter(filter map[string]interface{}) int64 {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := dClient.Collection(dao.Table).DeleteOne(ctx, filter)
	if err != nil {
		panic(err)
	}
	return result.DeletedCount
}

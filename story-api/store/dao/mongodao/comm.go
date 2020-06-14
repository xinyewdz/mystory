package mongodao

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
	"go.uber.org/zap"
	"story-api/global"
	"story-api/util/config"
	"time"
)

var(
	log = global.MainLog
	dClient *mongo.Database
)

func init(){
	host := config.Get("mongodb.host")
	database := config.Get("mongodb.database")
	user := config.Get("mongodb.user")
	password := config.Get("mongodb.password")
	url := "mongodb://"+user+":"+password+"@"+host+"/admin"
	ctx,_ := context.WithTimeout(context.Background(),5*time.Second)
	client,err := mongo.Connect(ctx,options.Client().ApplyURI(url))
	if err!=nil{
		log.Error("connect mongodao error.",zap.String("host",host),zap.Error(err))
	}
	dClient = client.Database(database)
}

func toDoc(v interface{})*bsoncore.Document{
	data,err := bson.Marshal(v)
	if err!=nil{
		return nil
	}
	doc := &bsoncore.Document{}
	bson.Unmarshal(data,doc)
	return doc
}

func fromDoc(doc *bsoncore.Document,v interface{}){
	data,err := bson.Marshal(doc)
	if err!=nil{
		return
	}
	err = bson.Unmarshal(data,v)
	if err!=nil{
		log.Error("fromDoc error",zap.Error(err))
	}
}


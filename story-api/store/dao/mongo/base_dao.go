package mongo
import (
	"github.com/syndtr/goleveldb/leveldb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
	"go.uber.org/zap"
	"story-api/global"
	"story-api/util/config"
)

var(
	log = global.MainLog
	host string
	database string
)

type daoType uint

const(
	User daoType = iota
	Story
)

func init(){
	host = config.Get("mongo.host")
	database = config.Get("mongo.database")
}

func getDb(dao daoType)*leveldb.DB{
	return nil
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


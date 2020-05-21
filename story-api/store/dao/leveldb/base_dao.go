package leveldb

import (
	"github.com/syndtr/goleveldb/leveldb"
	"go.uber.org/zap"
	"story-api/global"
)

var(
	dbMap = make(map[interface{}]*leveldb.DB)
	log = global.MainLog
)

func init(){
	userDb,err:= leveldb.OpenFile("db/user.db",nil)
	if err!=nil{
		log.Error("open userdb error",zap.Error(err))
	}
	storyDb,err:= leveldb.OpenFile("db/story.db",nil)
	if err!=nil{
		log.Error("open storydb error",zap.Error(err))
	}
	dbMap[&StoryDao{}]=storyDb
	dbMap[&UserDao{}]=userDb
}

func getDb(dao interface{})*leveldb.DB{
	return dbMap[dao]
}



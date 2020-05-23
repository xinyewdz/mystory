package leveldb

import (
	"github.com/syndtr/goleveldb/leveldb"
	"go.uber.org/zap"
	"story-api/global"
)

var(
	dbMap = make(map[daoType]*leveldb.DB)
	log = global.MainLog
)

type daoType uint

const(
	User daoType = iota
	Story
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
	dbMap[Story]=storyDb
	dbMap[User]=userDb
}

func getDb(dao daoType)*leveldb.DB{
	return dbMap[dao]
}



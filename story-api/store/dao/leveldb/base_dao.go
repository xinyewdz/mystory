package leveldb

import "github.com/syndtr/goleveldb/leveldb"

var(
	dbMap = make(map[interface{}]*leveldb.DB)
)

func init(){
	userDb,_:= leveldb.OpenFile("db/user.db",nil)
	storyDb,_ := leveldb.OpenFile("db/story.db",nil)
	dbMap[&StoryDao{}]=storyDb
	dbMap[&UserDao{}]=userDb
}

func getDb(dao interface{})*leveldb.DB{
	return dbMap[dao]
}



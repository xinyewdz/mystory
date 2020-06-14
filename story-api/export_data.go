package main

import (
	"story-api/store/dao/leveldb"
	"story-api/store/dao/mongodao"
)

func main(){
	storyDao := leveldb.NewStoryDao()
	mdao := mongodao.NewStoryDao()
	sl := storyDao.List()
	for _,story:= range sl{
		mdao.Insert(story)
	}
}

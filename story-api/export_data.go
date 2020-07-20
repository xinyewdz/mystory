package main

import (
	"story-api/store/dao/leveldb"
	"story-api/store/dao/mongodao"
)

func main() {
	convertStory()
}

func convertStory() {
	sdao := mongodao.NewStoryDao()
	sl := sdao.List(nil, "")
	for _, s := range sl {
		s.IsPublic = true
		sdao.UpdateObj(s.Id, s)
	}

}

func exportData() {
	storyDao := leveldb.NewStoryDao()
	mdao := mongodao.NewStoryDao()
	sl := storyDao.List()
	for _, story := range sl {
		mdao.Insert(story)
	}
}

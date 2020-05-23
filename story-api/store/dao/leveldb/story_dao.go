package leveldb

import (
	"encoding/json"
	"github.com/syndtr/goleveldb/leveldb"
	"story-api/store/entity"
	"strconv"
	"time"
)


var storyDb *leveldb.DB

type StoryDao struct {
}

func init(){
	storyDb = getDb(Story)
}

func (dao *StoryDao) Insert(obj *entity.DBStory){
	id := time.Now().Unix()
	obj.CreateTime = time.Now()
	obj.Id = id
	idStr := strconv.Itoa(int(obj.Id))
	sj,_ := json.Marshal(obj)
	storyDb.Put([]byte(idStr),sj,nil)
}

func (dao *StoryDao) Update(obj *entity.DBStory){
	idStr := strconv.Itoa(int(obj.Id))
	sJson,_ := json.Marshal(obj)
	storyDb.Put([]byte(idStr),sJson,nil)
}

func (dao *StoryDao)  List()[]*entity.DBStory{
	sl := []*entity.DBStory{}
	iterator := storyDb.NewIterator(nil,nil)
	for iterator.Next(){
		data := iterator.Value()
		obj := &entity.DBStory{}
		json.Unmarshal(data,obj)
		sl = append(sl,obj)
	}
	return sl
}

func (dao *StoryDao) Detail(id int64)*entity.DBStory{
	key := strconv.Itoa(int(id))
	valStr,_  := storyDb.Get([]byte(key),nil)
	s := &entity.DBStory{}
	json.Unmarshal(valStr,s)
	return s
}

func (dao *StoryDao) Remove(id int64){
	key := strconv.Itoa(int(id))
	storyDb.Delete([]byte(key),nil)
}

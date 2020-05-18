package leveldb

import (
	"encoding/json"
	"story-api/store/entity"
	"strconv"
	"time"
)


type StoryDao struct {
}

func (dao *StoryDao) Insert(obj *entity.DBStory){
	id := time.Now().Unix()
	obj.Id = id
	idStr := strconv.Itoa(int(obj.Id))
	sj,_ := json.Marshal(obj)
	getDb(dao).Put([]byte(idStr),sj,nil)
}

func (dao *StoryDao) UserUpdate(obj *entity.DBStory){
	idStr := strconv.Itoa(int(obj.Id))
	sJson,_ := json.Marshal(obj)
	getDb(dao).Put([]byte(idStr),sJson,nil)
}

func (dao *StoryDao)  List()[]*entity.DBStory{
	sl := []*entity.DBStory{}
	iterator := getDb(dao).NewIterator(nil,nil)
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
	valStr,_  := getDb(dao).Get([]byte(key),nil)
	s := &entity.DBStory{}
	json.Unmarshal(valStr,s)
	return s
}

func (dao *StoryDao) Remove(id int64){
	key := strconv.Itoa(int(id))
	getDb(dao).Delete([]byte(key),nil)
}
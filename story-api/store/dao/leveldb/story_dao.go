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

func NewStoryDao()*StoryDao{
	dao := &StoryDao{}
	return dao
}

func (dao *StoryDao) Insert(obj *entity.DBStory){
	id := time.Now().Unix()
	obj.CreateTime = time.Now()
	obj.Id = strconv.Itoa(int(id))
	sj,_ := json.Marshal(obj)
	storyDb.Put([]byte(obj.Id),sj,nil)
}

func (dao *StoryDao) Update(obj *entity.DBStory){
	sJson,_ := json.Marshal(obj)
	storyDb.Put([]byte(obj.Id),sJson,nil)
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

func (dao *StoryDao) Detail(id string)*entity.DBStory{
	valStr,_  := storyDb.Get([]byte(id),nil)
	s := &entity.DBStory{}
	json.Unmarshal(valStr,s)
	return s
}

func (dao *StoryDao) Remove(id string){
	storyDb.Delete([]byte(id),nil)
}

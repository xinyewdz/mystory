package leveldb

import (
	"encoding/json"
	"story-api/store/entity"
	"strconv"
	"time"
)

type UserDao struct {
}


func (dao *UserDao) Insert(obj *entity.DBUser){
	id := time.Now().Unix()
	obj.Id = id
	idStr := strconv.Itoa(int(obj.Id))
	sj,_ := json.Marshal(obj)
	getDb(dao).Put([]byte(idStr),sj,nil)
}

func (dao *UserDao) Update(obj *entity.DBUser){
	idStr := strconv.Itoa(int(obj.Id))
	sJson,_ := json.Marshal(obj)
	getDb(dao).Put([]byte(idStr),sJson,nil)
}

func (dao *UserDao) List()[]*entity.DBUser{
	sl := []*entity.DBUser{}
	iterator := getDb(dao).NewIterator(nil,nil)
	for iterator.Next(){
		data := iterator.Value()
		obj := &entity.DBUser{}
		json.Unmarshal(data,obj)
		sl = append(sl,obj)
	}
	return sl
}

func (dao *UserDao) Detail(id int64)*entity.DBUser{
	key := strconv.Itoa(int(id))
	valStr,_  := getDb(dao).Get([]byte(key),nil)
	s := &entity.DBUser{}
	json.Unmarshal(valStr,s)
	return s
}

func (dao *UserDao) Remove(id int64){
	key := strconv.Itoa(int(id))
	getDb(dao).Delete([]byte(key),nil)
}
package leveldb

import (
	"encoding/json"
	"github.com/syndtr/goleveldb/leveldb"
	"story-api/store/entity"
	"strconv"
	"time"
)

var userDb *leveldb.DB

type UserDao struct {
}

func init(){
	userDb = getDb(User)
}


func (dao *UserDao) Insert(obj *entity.DBUser){
	id := time.Now().Unix()
	idStr := strconv.Itoa(int(id))
	obj.Id = idStr
	sj,_ := json.Marshal(obj)
	userDb.Put([]byte(idStr),sj,nil)
}

func (dao *UserDao) Update(obj *entity.DBUser){
	sJson,_ := json.Marshal(obj)
	userDb.Put([]byte(obj.Id),sJson,nil)
}

func (dao *UserDao) List()[]*entity.DBUser{
	sl := []*entity.DBUser{}
	iterator := userDb.NewIterator(nil,nil)
	for iterator.Next(){
		data := iterator.Value()
		obj := &entity.DBUser{}
		json.Unmarshal(data,obj)
		sl = append(sl,obj)
	}
	return sl
}

func (dao *UserDao) Get(id int64)*entity.DBUser{
	key := strconv.Itoa(int(id))
	valStr,_  := userDb.Get([]byte(key),nil)
	s := &entity.DBUser{}
	json.Unmarshal(valStr,s)
	return s
}

func (dao *UserDao) GetByOpenId(openId string)*entity.DBUser{
	iterator := userDb.NewIterator(nil,nil)
	var user *entity.DBUser
	for iterator.Next(){
		data := iterator.Value()
		obj := &entity.DBUser{}
		json.Unmarshal(data,obj)
		if obj.Openid==openId{
			user = obj
			break
		}
	}
	return user
}

func (dao *UserDao) Remove(id int64){
	key := strconv.Itoa(int(id))
	userDb.Delete([]byte(key),nil)
}
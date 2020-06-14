package mongodao

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"story-api/store/entity"
	"time"
)

type UserDao struct {
	BaseDao
}

func NewUserDao()*UserDao{
	userdao := &UserDao{
	}
	userdao.Table = "user"
	userdao.Obj = entity.DBUser{}
	return userdao
}

func (dao *UserDao) Insert(obj *entity.DBUser){
	obj.CreateTime = time.Now()
	obj.Id = primitive.NewObjectID().Hex()
	dao.InsertObj(obj)
}

func (dao *UserDao) Update(obj *entity.DBUser){
 	id := obj.Id
 	dao.UpdateObj(id,obj)
}

func (dao *UserDao) List()[]*entity.DBUser{
	sl := []*entity.DBUser{}
	list := dao.ListAll()
	if list==nil{
		return nil
	}
	for _,obj := range list{
		sl = append(sl,obj.(*entity.DBUser))
	}
	return sl
}

func (dao *UserDao) Get(id string)*entity.DBUser{
	return dao.GetObj(id).(*entity.DBUser)
}

func (dao *UserDao) GetByName(name string)*entity.DBUser{
	query := bson.M{
		"name":name,
	}
	ctx,_ := context.WithTimeout(context.Background(),5*time.Second)
	result := dClient.Collection(dao.Table).FindOne(ctx,query)
	obj := &entity.DBUser{}
	err := result.Decode(obj)
	if err==nil{
		return obj
	}
	if err==mongo.ErrNoDocuments{
		return nil
	}else{
		panic(err)
	}
}

func (dao *UserDao) GetByOpenId(openId string)*entity.DBUser{
	query := bson.M{
		"openId":openId,
	}
	ctx,_ := context.WithTimeout(context.Background(),5*time.Second)
	result := dClient.Collection(dao.Table).FindOne(ctx,query)
	obj := &entity.DBUser{}
	err := result.Decode(obj)
	if err==nil{
		return obj
	}
	if err==mongo.ErrNoDocuments{
		return nil
	}else{
		panic(err)
	}
}

func (dao *UserDao) Remove(id string){
	dao.RemoveObj(id)
}
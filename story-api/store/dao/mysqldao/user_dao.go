package mysqldao

import (
	"story-api/store/entity"
	"strconv"
	"time"
)

var(
	userInsert = "insert into user(open_id,name,gender,phone,password,type,create_time) values(?,?,?,?)"
	userUpdate  = "update user set open_id=?,name=?,gender=?,phone=?,password=?,type=? where id=?"
	userList    = "select * from user"
	userGet     = "select * from user where id=?"
	userDel     = "delete from user where id=?"
)

type UserDao struct {
	table string
}

func NewUserDao()*UserDao{
	dao := &UserDao{
		table: "user",
	}
	return dao
}

func (dao *UserDao) Insert(obj *entity.DBUser){
	obj.CreateTime = time.Now()
	result, err := conn.Exec(userInsert,obj.Openid,obj.Name,obj.Gender,obj.Phone,obj.Password,obj.Type,obj.CreateTime)
	if err!=nil{
		panic(err)
	}
	id,_ :=result.LastInsertId()
	obj.Id = strconv.Itoa(int(id))
}

func (dao *UserDao) Update(obj *entity.DBUser){
	_,err := conn.Exec(userUpdate,obj.Openid,obj.Name,obj.Gender,obj.Phone,obj.Password,obj.Type,obj.CreateTime,obj.Id)
	if err!=nil{
		panic(err)
	}
}

func (dao *UserDao)  List()[]*entity.DBUser{
	rows,err :=conn.Query(userList)
	if err!=nil{
		panic(err)
	}
	defer rows.Close()
	list := []*entity.DBUser{}
	for rows.Next(){
		obj := &entity.DBUser{}
		err = fromRows(rows,obj)
		if err!=nil{
			panic(err)
			return nil
		}
		list = append(list,obj)
	}
	return list
}

func (dao *UserDao) Detail(id string)*entity.DBUser{
	s := &entity.DBUser{}
	result,err := conn.Query(userGet,id)
	if result.Next(){
		err = fromRows(result,s)
	}
	if err!=nil{
		panic(err)
		return nil
	}
	return s
}

func (dao *UserDao) Remove(id string){
	_,err := conn.Exec(userDel,id)
	if err!=nil{
		panic(err)
	}
}

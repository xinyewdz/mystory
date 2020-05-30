package mysql

import (
	"story-api/store/entity"
	"strconv"
	"time"
)

var(
	table       = "story"
	storyInsert = "insert into story(name,audio_url,image_url,create_time) values(?,?,?,?)"
	storyUpdate  = "update story set name=?,audio_url=?,image_url=? where id=?"
	storyList    = "select * from story"
	storyGet     = "select audio_url,id,name,image_url,create_time from story where id=?"
	storyDel     = "delete from story where id=?"
)

type StoryDao struct {
	table string
}

func NewStoryDao()*StoryDao{
	dao := &StoryDao{
		table: "story",
	}
	return dao
}

func (dao *StoryDao) Insert(obj *entity.DBStory){
	obj.CreateTime = time.Now()
	id,err := insert(conn,dao.table,obj)
	if err!=nil{
		panic(err)
	}
	obj.Id = strconv.Itoa(int(id))
}

func (dao *StoryDao) Update(obj *entity.DBStory){
	_,err := conn.Exec(storyUpdate,obj.Name,obj.AudioUrl,obj.ImageUrl,obj.Id)
	if err!=nil{
		panic(err)
	}
}

func (dao *StoryDao)  List()[]*entity.DBStory{
	rows,err :=conn.Query(storyList)
	if err!=nil{
		panic(err)
	}
	defer rows.Close()
	list := []*entity.DBStory{}
	for rows.Next(){
		obj := &entity.DBStory{}
		err = fromRows(rows,obj)
		if err!=nil{
			panic(err)
			return nil
		}
		list = append(list,obj)
	}
	return list
}

func (dao *StoryDao) Detail(id string)*entity.DBStory{
	s := &entity.DBStory{}
	result,err := conn.Query(storyGet,id)
	if result.Next(){
		err = fromRows(result,s)
	}
	if err!=nil{
		panic(err)
		return nil
	}
	return s
}

func (dao *StoryDao) Remove(id string){
	_,err := conn.Exec(storyDel,id)
	if err!=nil{
		panic(err)
	}
}

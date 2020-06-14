package mongodao

import (
	"story-api/store/entity"
	"testing"
)

func TestInsertUser(t *testing.T){
	user := &entity.DBUser{
		Name: "aaron",
		Type: "admin",
	}
	userDao := NewUserDao()
	userDao.Insert(user)
}

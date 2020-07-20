package mongodao

import (
	"story-api/store/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StoryPlayDetailDao struct {
	BaseDao
}

func NewStoryPlayDetailDao() *StoryPlayDetailDao {
	dao := &StoryPlayDetailDao{}
	dao.Table = "story_play_detail"
	dao.Obj = entity.DBStoryPlayDetail{}
	return dao
}

func (dao *StoryPlayDetailDao) Insert(obj *entity.DBStoryPlayDetail) {
	obj.CreateTime = time.Now()
	obj.Id = primitive.NewObjectID().Hex()
	dao.InsertObj(obj)
}

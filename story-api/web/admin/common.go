package admin

import (
	"story-api/global"
	"story-api/store/dao/mongodao"
	"story-api/util/redisutil"

	"github.com/go-redis/redis/v8"
)

var (
	mainLog            = global.MainLog
	userDao            = mongodao.NewUserDao()
	storyDao           = mongodao.NewStoryDao()
	storyPlayDetailDao = mongodao.NewStoryPlayDetailDao()
	storyFavoriteDao   = mongodao.NewStoryFavoriteDao()
	redisClient        *redis.Client
)

func init() {
	redisClient = redisutil.Client
}

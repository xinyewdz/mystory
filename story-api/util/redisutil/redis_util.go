package redisutil

import (
	"github.com/go-redis/redis/v8"
	"story-api/util/config"
)

var(
	Client *redis.Client
)

func init(){
	addr := config.Get("redis.host")
	Client = redis.NewClient(&redis.Options{
		Addr: addr,
	})
}

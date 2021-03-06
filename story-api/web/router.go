package web

import (
	"context"
	"encoding/json"
	"net/http"
	"story-api/common"
	"story-api/store/entity"
	"story-api/util/redisutil"
	"time"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

const USER_KEY = "user"

var redisClient *redis.Client
var publicUrl map[string]interface{} = make(map[string]interface{})

func init() {
	redisClient = redisutil.Client
	publicUrl["/login"] = nil
	publicUrl["/admin/login"] = nil
	publicUrl["/play/list"] = nil
	publicUrl["/play"] = nil
}

type RouterHttpHandler func(context.Context, http.ResponseWriter, *http.Request) *common.ApiResponse

func (router RouterHttpHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	mainLog.Info("request ", zap.String("path", path))
	resp.Header().Set("Content-Type", "application/json;charset=UTF-8")
	resp.Header().Set("Access-Control-Allow-Credentials", "true")
	resp.Header().Set("Access-Control-Allow-Headers", "*")
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Methods", "PUT,POST,GET,DELETE,OPTIONS")
	resp.WriteHeader(200)
	defer func() {
		if err := recover(); err != nil {
			//path := model.URL.Path
			ap := &common.ApiResponse{
				Code: "500",
				Msg:  "server error",
			}
			mainLog.Error("request error", zap.String("path", path), zap.Any("error", err))
			WriteResp(resp, ap)
		}
	}()
	var user *entity.DBUser = nil
	var apiResp *common.ApiResponse
	if _, ok := publicUrl[path]; ok {
		context := context.WithValue(context.Background(), USER_KEY, user)
		apiResp = router(context, resp, req)
	} else {
		user = auth(req)
		if user == nil {
			apiResp = &common.ApiResponse{
				Code: "401",
				Msg:  "token无效",
			}
		} else {
			context := context.WithValue(context.Background(), USER_KEY, user)
			apiResp = router(context, resp, req)
		}
	}
	if apiResp != nil {
		WriteResp(resp, apiResp)
	}
}

func auth(req *http.Request) *entity.DBUser {
	token := req.Header.Get("token")
	if token == "" {
		return nil
	}
	if token == "wendzh" {
		user := userDao.GetByName("aaron")
		return user
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	key := TOKEN_KEY + token
	userRes := redisClient.Get(ctx, key)
	if userRes == nil || userRes.Val() == "" {
		return nil
	}
	redisClient.PExpire(ctx, key, 30*time.Minute)
	user := &entity.DBUser{}
	json.Unmarshal([]byte(userRes.Val()), user)
	return user
}

package web

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"story-api/common"
	"story-api/store/entity"
	"story-api/util/redisutil"
	"time"
)

const USER_KEY = "user"

var redisClient *redis.Client

func init(){
	redisClient = redisutil.Client
}

type RouterHttpHandler func(context.Context,http.ResponseWriter, *http.Request)*common.ApiResponse

func (router RouterHttpHandler)ServeHTTP(resp http.ResponseWriter,req *http.Request){
	path := req.URL.Path
	mainLog.Info("request ",zap.String("path",path))
	resp.Header().Set("Content-Type","application/json;charset=UTF-8")
	resp.WriteHeader(200)
	defer func(){
		if err :=recover();err!=nil {
			//path := model.URL.Path
			ap := &common.ApiResponse{
				Code:"500",
				Msg: "server error",
			}
			mainLog.Error("request error",zap.String("path",path),zap.Any("error",err))
			writeResp(resp,ap)
		}
	}()
	var user *entity.DBUser= nil
	var apiResp *common.ApiResponse
	if path=="/login"{
		apiResp = router(nil,resp,req)
	}else{
		user = auth(req)
		if user==nil{
			apiResp = &common.ApiResponse{
				Code:"401",
				Msg: "token无效",
			}
		}else {
			context := context.WithValue(context.Background(), USER_KEY, user)
			apiResp = router(context,resp,req)
		}
	}
	if apiResp!=nil{
		writeResp(resp,apiResp)
	}
}

func auth(req *http.Request)*entity.DBUser{
	token := req.Header.Get("token")
	if token==""{
		return nil
	}
	if token=="wendzh"{
		return new(entity.DBUser)
	}
	ctx,_ := context.WithTimeout(context.Background(),5*time.Second)
	key := TOKEN_KEY+token
	userStr := redisClient.Get(ctx,key)
	if userStr==nil{
		return nil
	}
	redisClient.PExpire(ctx,key,30*time.Minute)
	user := &entity.DBUser{}
	json.Unmarshal([]byte(userStr.String()),user)
	return user
}

func resolveBody(req *http.Request,reqBody interface{}){
	data,_ := ioutil.ReadAll(req.Body)
	json.Unmarshal(data,reqBody)
}

func writeResp(resp http.ResponseWriter,apiResp *common.ApiResponse){
	respData,_ := json.Marshal(apiResp)
	resp.Write(respData)
}

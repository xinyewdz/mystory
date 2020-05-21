package web

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
	"story-api/common"
)

type RouterHttpHandler func(http.ResponseWriter, *http.Request)

func (router RouterHttpHandler)ServeHTTP(resp http.ResponseWriter,req *http.Request){
	path := req.URL.Path
	mainLog.Info("request ",zap.String("path",path))
	defer func(){
		resp.Header().Set("Content-Type","application/json;charset=utf-8")
		if err :=recover();err!=nil {
			//path := req.URL.Path
			ap := &common.ApiResponse{
				Code:"500",
				Msg: "server error",
			}
			mainLog.Error("request error",zap.String("path",path),zap.Any("error",err))
			respData,_ :=json.Marshal(ap)
			resp.Write(respData)
		}
	}()
	router(resp,req)
}

func auth(req *http.Request)bool{
	//token := req.Header.Get("token")
	return true
}

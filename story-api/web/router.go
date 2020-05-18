package web

import (
	"encoding/json"
	"log"
	"net/http"
	"story-api/common"
)

type RouterHttpHandler func(http.ResponseWriter, *http.Request)

func (router RouterHttpHandler)ServeHTTP(resp http.ResponseWriter,req *http.Request){
	path := req.URL.Path
	log.Printf("request %s\n",path)
	defer func(){
		resp.Header().Set("Content-Type","application/json;charset=utf-8")
		if err :=recover();err!=nil {
			//path := req.URL.Path
			ap := &common.ApiResponse{
				Code:"500",
				Msg: "server error",
			}
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

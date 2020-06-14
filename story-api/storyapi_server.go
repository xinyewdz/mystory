package main

import (
	"log"
	"net/http"
	"story-api/web"
)


var(
	storyWeb = new(web.StoryWeb)
	userWeb = new(web.UserWeb)
)

func init(){
}

func main(){
	mux := http.NewServeMux()
	regist(mux)
	port := "8060"
	log.Printf("server start at:%s\n",port)
	http.ListenAndServe(":"+port,mux)
}

func regist(mux *http.ServeMux){
	routeMap := make(map[string]web.RouterHttpHandler)
	registStory(routeMap)
	registUser(routeMap)
	for k,v := range routeMap{
		mux.Handle(k,v)
	}
}

func registStory(routeMap map[string]web.RouterHttpHandler){
	routeMap["/list"] = storyWeb.List
	routeMap["/story"] = storyWeb.Detail
	routeMap["/upload"] = storyWeb.Upload
	routeMap["/save"] = storyWeb.Save
	routeMap["/remove"] = storyWeb.Remove

}

func registUser(routeMap map[string]web.RouterHttpHandler){
	routeMap["/user/detail"] = userWeb.Detail
	routeMap["/user/list"] = userWeb.List
	routeMap["/user/save"] = userWeb.Save
	routeMap["/user/remove"] = userWeb.Remove
	routeMap["/login"] = userWeb.Login
}





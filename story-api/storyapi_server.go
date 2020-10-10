package main

import (
	"flag"
	"log"
	"net/http"
	"story-api/web"
	"story-api/web/admin"
)

var (
	storyWeb             = new(web.StoryWeb)
	userWeb              = new(web.UserWeb)
	userAdminWeb         = new(admin.UserAdminWeb)
	enableStoryList bool = true
	help                 = false
)

func init() {
	flag.BoolVar(&enableStoryList, "sl", true, "enable story list")
	flag.BoolVar(&help, "h", false, "help usage")
	flag.Parse()
}

func main() {
	if help {
		flag.PrintDefaults()
		return
	}
	userWeb.EnableStoryList = enableStoryList
	log.Println("server run enableStoryList", enableStoryList)
	mux := http.NewServeMux()
	regist(mux)
	port := "8060"
	log.Printf("server start at:%s\n", port)
	http.ListenAndServe(":"+port, mux)
}

func regist(mux *http.ServeMux) {
	routeMap := make(map[string]web.RouterHttpHandler)
	registryStory(routeMap)
	registryUser(routeMap)
	registryAdmin(routeMap)
	for k, v := range routeMap {
		mux.Handle(k, v)
	}
}

func registryAdmin(routeMap map[string]web.RouterHttpHandler) {
	routeMap["/admin/login"] = userAdminWeb.Login
	routeMap["/admin/story/playDetail"] = userAdminWeb.ListStoryPlayDetail
	routeMap["/admin/user/playDetail"] = userAdminWeb.ListUserPlayDetail
}

func registryStory(routeMap map[string]web.RouterHttpHandler) {
	routeMap["/play/list"] = storyWeb.PlayList
	routeMap["/list"] = storyWeb.List
	routeMap["/story"] = storyWeb.Detail
	routeMap["/upload"] = storyWeb.Upload
	routeMap["/save"] = storyWeb.Save
	routeMap["/remove"] = storyWeb.Remove
	routeMap["/play"] = storyWeb.Play

}

func registryUser(routeMap map[string]web.RouterHttpHandler) {
	routeMap["/user/detail"] = userWeb.Detail
	routeMap["/user/list"] = userWeb.List
	routeMap["/user/save"] = userWeb.Save
	routeMap["/user/remove"] = userWeb.Remove
	routeMap["/login"] = userWeb.Login
}

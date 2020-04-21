package main

import (
	"bufio"
	"encoding/json"
	"github.com/syndtr/goleveldb/leveldb"
	"io"
	"log"
	"net/http"
	"os"
	"story-api/common"
	"strconv"
	"strings"
)

type Story struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Image string `json:"image"`
	Url string `json:"url"`
}

var(
	ld *leveldb.DB
	sl []Story
)

func init(){
	ld,_= leveldb.OpenFile("conf/story.db",nil)
	sl = getStoryList("conf/story.conf")
	for _,s := range sl{
		skey := strconv.Itoa(int(s.Id))
		data,_ := json.Marshal(s)
		ld.Put([]byte(skey),data,nil)
	}
}

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/list",http.HandlerFunc(list))
	mux.HandleFunc("/story",http.HandlerFunc(story))
	http.ListenAndServe(":8005",mux)
}

func list(resp http.ResponseWriter,req *http.Request){
	ar := common.ApiResponse{}
	ar.Code = "200"
	ar.Data = sl
	data,_ := json.Marshal(ar)
	resp.Write(data)
}

func story(resp http.ResponseWriter,req *http.Request){
	req.ParseForm()
	idStr := req.Form.Get("id")
	id,_ := strconv.Atoi(idStr)
	key := strconv.Itoa(id)
	log.Println(idStr)
	valStr,_ := ld.Get([]byte(key),nil)
	s := Story{}
	json.Unmarshal(valStr,&s)
	ar := common.ApiResponse{
		Data: &s,
	}
	if valStr!=nil {
		ar.Success(s)
	}
	data,_ := json.Marshal(ar)
	resp.Write(data)
}

func getStoryList(path string)[]Story{
	f,_ := os.Open(path)
	reader := bufio.NewReader(f)
	storyList := []Story{}
	for{
		line,_,err := reader.ReadLine()
		if err==io.EOF {
			break
		}
		linStr := string(line)
		lis := strings.Split(linStr,",")
		s := Story{}
		s.Id,_ = strconv.ParseInt(lis[0],10,0)
		s.Name = lis[1]
		s.Image = lis[2]
		s.Url = lis[3]
		storyList = append(storyList,s)
	}
	return storyList
}

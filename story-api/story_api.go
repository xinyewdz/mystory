package main

import (
	"bufio"
	"encoding/json"
	"github.com/syndtr/goleveldb/leveldb"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"story-api/common"
	"story-api/util"
	"strconv"
	"strings"
	"time"
)

type Story struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Image string `json:"image"`
	Url string `json:"url"`
}

var(
	ld *leveldb.DB
)

func init(){
	ld,_= leveldb.OpenFile("conf/story.db",nil)
}

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/list",http.HandlerFunc(list))
	mux.HandleFunc("/story",http.HandlerFunc(detail))
	mux.HandleFunc("/upload",http.HandlerFunc(upload))
	mux.HandleFunc("/save",http.HandlerFunc(saveStory))
	mux.HandleFunc("/remove",http.HandlerFunc(remove))
	http.ListenAndServe(":8060",mux)
}

func upload(resp http.ResponseWriter,req *http.Request){
	req.ParseMultipartForm(1024*1024*100)
	file,fh,_ := req.FormFile("name")
	fileName := fh.Filename
	name := req.Form.Get("name")
	fileType := fileName[strings.LastIndex(fileName,".")+1:]
	filePath := util.UpyunUpload(file,name+"."+fileType)
	ar := &common.ApiResponse{
		Data:"",
	}
	ar.Success(filePath)
	data,_ := json.Marshal(ar)
	resp.Write(data)
}

func saveStory(resp http.ResponseWriter,req *http.Request){
	//req.ParseMultipartForm(1024*1024*10)
	data,_ := ioutil.ReadAll(req.Body)
	dataMap := make(map[string]string)
	json.Unmarshal(data,&dataMap)
	name := dataMap["name"]
	audio := dataMap["audio"]
	image := dataMap["image"]
	story := &Story{}
	story.Name = name
	story.Image = image
	story.Url = audio
	story.Id = time.Now().Unix()
	idStr := strconv.Itoa(int(story.Id))
	sj,_ := json.Marshal(story)
	ld.Put([]byte(idStr),sj,nil)
	ar := &common.ApiResponse{
		Data:0,
	}
	ar.Success(story.Id)
	respData,_ := json.Marshal(ar)
	resp.Write(respData)
}



func list(resp http.ResponseWriter,req *http.Request){
	ar := &common.ApiResponse{
		Data:&[]Story{},
	}
	sl := []Story{}
	iterator := ld.NewIterator(nil,nil)
	for iterator.Next(){
		data := iterator.Value()
		story := &Story{}
		json.Unmarshal(data,&story)
		sl = append(sl,*story)
	}
	ar.Success(sl)
	data,_ := json.Marshal(ar)
	resp.Write(data)
}

func remove(resp http.ResponseWriter,req *http.Request){
	req.ParseForm()
	idStr := req.Form.Get("id")
	id,_ := strconv.Atoi(idStr)
	key := strconv.Itoa(id)
	log.Println("remove:"+idStr)
	ld.Delete([]byte(key),nil)
	ar := &common.ApiResponse{
	}
	ar.Success(nil)
	data,_ := json.Marshal(ar)
	resp.Write(data)
}

func detail(resp http.ResponseWriter,req *http.Request){
	req.ParseForm()
	idStr := req.Form.Get("id")
	id,_ := strconv.Atoi(idStr)
	key := strconv.Itoa(id)
	log.Println(idStr)
	valStr,_ := ld.Get([]byte(key),nil)
	log.Println(string(valStr))
	s := Story{}
	json.Unmarshal(valStr,&s)
	ar := &common.ApiResponse{

	}
	ar.Success(s)
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

package util

import (
	"bufio"
	"io"
	"os"
	"strings"
)

var confMap = map[string]string{}

func init(){
	f,_ := os.Open("conf/story.conf")
	reader := bufio.NewReader(f)
	for{
		line,_,err := reader.ReadLine()
		if err==io.EOF {
			break
		}
		linStr := string(line)
		datas := strings.Split(linStr,"=")
		confMap[datas[0]] = datas[1]
	}

}


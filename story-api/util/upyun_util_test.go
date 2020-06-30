package util

import (
	"os"
	"testing"
)

func TestUpyunUpload(t *testing.T) {
	f,_ := os.Open("upyun_util.go")
	uPath := UpyunUpload(f,"test")
	t.Log(uPath)
}

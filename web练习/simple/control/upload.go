package control

import (
	"encoding/json"
	"io"
	"math/rand"
	"path"
	"time"
	"unsafe"

	"net/http"
	"os"
)

func ApiUpload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1 << 20 * 20) //准备一个容器，20MB空间
	f, h, err := r.FormFile("upfile")  //通过找对应的文件
	if err != nil {
		Fail(w, "上传失败", err.Error())
	}
	os.MkdirAll("static/upload/", 0666) //调用，然后给权限。任何人都可以读写
	ext := path.Ext(h.Filename)
	name := "static/upload/" + RandStr(10)+ext
	dst, _ := os.Create(name) //指定路径
	io.Copy(dst, f)
	f.Close()
	dst.Close()
	//Succ(w,"上传成功","/"+name)//xxx是路径
	w.Header().Set("Content-Type", "application/json")
	//富文本编辑器约定的格式
	mod := editorReply{
		Original: h.Filename,
		State:    "SUCCESS",
		Title:    h.Filename,
		Url:      "/" + name,
	}
	w.Write(mod.Json())
}

//全局种子
func init(){
	rand.Seed(time.Now().UnixNano())
}

var all="abcdefghijklmnopqrstuvwxyz0123456789"

//构建随机字符串
func RandStr(ln int)string {
	//rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	res := make([]byte, 0, ln)
	for i := 0; i < ln; i++ {
		res[i] = all[rand.Intn(36)]
	}
	return *(*string)(unsafe.Pointer(&res))
}

type editorReply struct {
	Original string `json:"original"`
	State    string `json:"state"`
	Title    string `json:"title"`
	Url      string `json:"url"`
}

//方法
func (er *editorReply) Json() []byte {
	buf, _ := json.Marshal(er)
	return buf
}

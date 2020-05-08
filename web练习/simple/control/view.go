package control

import (
	"io"
	"net/http"
	"os"
)

//主页面操作
func ArticleIndex(w http.ResponseWriter,r *http.Request){
	f, _ := os.Open("./view/index.html")
	io.Copy(w,f)
	f.Close()
}

//list列表页
func ListView(w http.ResponseWriter,r *http.Request){
	f, _ := os.Open("./view/list.html")
	io.Copy(w,f)
	f.Close()
}

//文章添加页面
func ADDView(w http.ResponseWriter,r *http.Request){
	f, _ := os.Open("./view/add.html")
	io.Copy(w,f)
	f.Close()
}

//详细页面
func DetailView(w http.ResponseWriter,r *http.Request){
	f, _ := os.Open("./view/detail.html")
	io.Copy(w,f)
	f.Close()
}

//修改页面
func EditView(w http.ResponseWriter,r *http.Request){
	f, _ := os.Open("./view/edit.html")
	io.Copy(w,f)
	f.Close()
}

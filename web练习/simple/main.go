package main

import (
	"net/http"
	"simple/control"
)

func main(){
	//设置静态目录
	http.Handle("/static/",http.StripPrefix("/static/",http.FileServer(http.Dir("static/"))))
	http.HandleFunc("/",control.ArticleIndex)
	http.HandleFunc("/api/index/data",control.ArticleData)

	http.HandleFunc("/list",control.ListView)//页面
	http.HandleFunc("/api/list/data",control.ArticleListData)
	http.HandleFunc("/api/list/del",control.ArticleListDel)

	http.HandleFunc("/add",control.ADDView)//添加页面
	http.HandleFunc("/api/article/add",control.ArticleADD)//添加

	http.HandleFunc("/detail",control.DetailView)//详细页面

	http.HandleFunc("/edit",control.EditView)//修改页面
	http.HandleFunc("/api/article/edit",control.ArticleEdit)//编辑

	http.HandleFunc("/api/upload",control.ApiUpload)
	http.ListenAndServe(":8086",nil)
}

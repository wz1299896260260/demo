package control

import (
	"net/http"
	"simple/model"
)



//list  select
func ArticleListData(w http.ResponseWriter,r *http.Request){
	mods, err := model.ArticleList()
	if err!=nil{
		Fail(w,err.Error())
		return
	}
	Succ(w,"列表",mods)
}


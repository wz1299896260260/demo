package control

import (
	"net/http"
	"simple/model"
	"strconv"
)

//del 删除
func ArticleListDel(w http.ResponseWriter,r *http.Request){
	r.ParseForm()//解析
	//w.Header().Set("Content-Type","application/json")
	idstr := r.Form.Get("id")
	id, _ := strconv.ParseInt(idstr, 10, 64)
	if model.ArticleDel(id){
		Succ(w,"删除成功")
		return
	}
	Fail(w,"删除失败")
	return
}
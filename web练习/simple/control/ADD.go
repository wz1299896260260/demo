package control

import (
	"encoding/json"
	"net/http"
	"simple/model"
	"time"
)

func ArticleADD(w http.ResponseWriter, r *http.Request) {
	/*r.ParseForm()
	mod := &model.Article{}
	mod.Title = r.Form.Get("title")
	mod.Author = r.Form["author"][0]
	mod.Content = r.FormValue("content")
	mod.Hits, _ = strconv.Atoi(r.Form.Get("hits"))
	mod.Utime = time.Now()
	log.Println(mod)*/
	//payload格式
	mod := &model.Article{}
	err := json.NewDecoder(r.Body).Decode(mod)//创建解码器解码。NewDecoder有个io.Reader接口
	if err != nil {
		Fail(w, "输入数据有误",err.Error())
		return
	}
	mod.Utime=time.Now()
	err = model.ArticleAdd(mod)
	if err != nil {
		Fail(w, "添加失败",err.Error())
		return
	}
	Succ(w, "添加成功")
	return

}



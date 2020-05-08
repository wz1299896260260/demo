package control

import (
	"encoding/json"
	"net/http"
	"simple/model"
	"strconv"
)

//首页数据
func ArticleData(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析
	w.Header().Set("Content-Type", "application/json")
	idstr := r.Form.Get("id")
	id, _ := strconv.ParseInt(idstr, 10, 64)
	mod, err := model.ArticleGet(id)
	if err != nil {
		Fail(w, err.Error())
		return
	}
	buf, _ := json.Marshal(mod)
	w.Write(buf)
}

func Fail(w http.ResponseWriter, msg string,data ...interface{}) {
	mod := Reply{
		Code: 300,//失败
		Msg:  msg,
	}
	if len(data)>0{
		mod.Data=data[0]//可传可不传
	}
	buf, _ := json.Marshal(mod)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(buf))
}

func Succ(w http.ResponseWriter, msg string,data ...interface{}) {//...interface{}可变参数，可以传也可以不传
	mod := Reply{
		Code: 200,
		Msg:  msg,
	}
	if len(data)>0{
		mod.Data=data[0]//可传可不传
	}
	buf, _ := json.Marshal(mod)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(buf))
}

type Reply struct {
	Code int         `json:"code"`//状态码
	Msg  string      `json:"msg"`//给用户信息
	Data interface{} `json:"data"`//返回数据/返回开发者查看的错误信息
}

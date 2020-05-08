package control

import (
	"encoding/json"
	"net/http"
	"simple/model"
)

func ArticleEdit(w http.ResponseWriter, r *http.Request) {
	//payload格式
	mod := &model.Article{}
	err := json.NewDecoder(r.Body).Decode(mod)//创建解码器解码。NewDecoder有个io.Reader接口
	if err != nil {
		Fail(w, "输入数据有误",err.Error())
		return
	}

	err = model.ArticleEdit(mod)
	if err != nil {
		Fail(w, "修改失败",err.Error())
		return
	}
	Succ(w, "修改成功")
	return

}

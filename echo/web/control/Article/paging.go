package Article

import (
	"github.com/labstack/echo"
	"github.com/zxysilent/utils"

	"web/control"
	"web/model"
)

func ArticlePage(e echo.Context)error{
	ipt:=&control.Page{}
	err := e.Bind(ipt)
	if err!=nil{
		return e.JSON(utils.ErrIpt("输入数据有误",err.Error()))//ErrIpt数据
	}
	if ipt.Pi<1{
		ipt.Pi=1
	}
	if ipt.Ps<1||ipt.Ps>50{
		ipt.Ps=10000 //超过范围给默认值
	}
	count := model.ArticleCount()
	if count<1{
		return e.JSON(utils.ErrOpt("未查询到数据"))
	}
	mods, err := model.ArticlePage(ipt.Pi,ipt.Ps)
	if err!=nil{
		return e.JSON(utils.ErrOpt("未查询到数据",err.Error()))
	}
	return e.JSON(utils.Page("评论数据",mods,count))
}

package Article

import (
	"github.com/labstack/echo"
	"github.com/zxysilent/utils"
	"web/model"
)

func ArticleAdd(e echo.Context)error{
	ipt := &model.Article{}
	err := e.Bind(ipt)
	if err!=nil{
		return e.JSON(utils.ErrIpt("输入数据有误",err.Error()))
	}
	err = model.ArticleAdd(ipt)

	if err!=nil{
		return e.JSON(utils.Fail("添加失败",err.Error()))
	}
	return e.JSON(utils.Succ("添加成功"))

}

package Article

import (
	"github.com/labstack/echo"
	"github.com/zxysilent/utils"
	"web/model"
)

func ArticleAll(e echo.Context)error{
	mods, err := model.ArticleAll()
	if err!=nil{
		return e.JSON(utils.Fail("未查询到数据",err.Error()))
	}

	return e.JSON(utils.Succ("Article数据",mods))
}

package Article

import (
	"github.com/labstack/echo"
	"github.com/zxysilent/utils"
	"strconv"
	"web/model"
)

func ArticleDel(e echo.Context)error{
	uid, err := strconv.ParseInt(e.Param("uid"), 10, 64)
	if err!=nil {
		return e.JSON(utils.ErrIpt("输入数据有误", err.Error()))
	}
	err = model.ArticlerDel(uid)
	if err!=nil {
		return e.JSON(utils.Fail("删除失败", err.Error()))
	}
	return e.JSON(utils.Succ("删除成功"))
}

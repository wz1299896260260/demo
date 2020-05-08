package Particular


import (
	"github.com/labstack/echo"
	"github.com/zxysilent/utils"
	"strconv"
	"web/model"
)

func ParticularDel(e echo.Context)error{
	id, err := strconv.ParseInt(e.Param("id"), 10, 64)
	if err!=nil {
		return e.JSON(utils.ErrIpt("输入数据有误", err.Error()))
	}
	err = model.ParticularDel(id)
	if err!=nil {
		return e.JSON(utils.Fail("删除失败", err.Error()))
	}
	return e.JSON(utils.Succ("删除成功"))
}

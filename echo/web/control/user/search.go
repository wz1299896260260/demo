package user

import (
	"github.com/labstack/echo"
	"github.com/zxysilent/utils"

	"web/model"
)

func UserGet(e echo.Context)error{
	//id, err := strconv.ParseInt(e.Param("id"), 0, 64)
	username := e.Param("username")

	//if err!=nil{
	//	return e.JSON(utils.ErrIpt("输入有误",err.Error()))
	//}
	mod, err := model.UserGet(username)
	if err!=nil{
		return e.JSON(utils.ErrOpt("未查询到数据",err.Error()))
	}
	return e.JSON(utils.Succ("用户数据",mod))
}


package Particular

import (
	
	"github.com/labstack/echo"
	"github.com/zxysilent/utils"


	"web/model"
)

//学校
func ParticularSchoolSearch(e echo.Context)error{
	school := e.Param("school") //获取path参数
	//class := e.Request().URL.Query()
	//fmt.Println(class) 获取参数
	//id, err := strconv.ParseInt(e.Param("id"), 10, 64)
	//if err!=nil {
	//	return e.JSON(utils.ErrIpt("输入数据有误", err.Error()))
	//}
	mod, err := model.ParticularSchoolSearchGet(school)
	if err!=nil {
		return e.JSON(utils.ErrOpt("未查询到数据", err.Error()))//ErrOpt没有数据返回
	}
	return e.JSON(utils.Succ("Particular数据",mod))
}

//班级
func ParticularClassSearch(e echo.Context)error{
	class := e.Param("class") //获取path参数
	mod, err := model.ParticularClassSearchGet(class)
	if err!=nil {
		return e.JSON(utils.ErrOpt("未查询到数据", err.Error()))//ErrOpt没有数据返回
	}
	return e.JSON(utils.Succ("Particular数据",mod))
}
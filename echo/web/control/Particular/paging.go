package Particular

import (
	"github.com/labstack/echo"
	"github.com/zxysilent/utils"
	"strconv"
	"web/model"
)

//分页和数据总数一起是查询所有分类
func ParticularALL(e echo.Context)error{
	mods, err := model.ParticularAll()
	if err!=nil{
		return e.JSON(utils.Fail("未查询到数据",err.Error()))
	}

	return e.JSON(utils.Succ("Particular数据",mods))
}

//分页查询
func ParticularPage(e echo.Context)error{
	pi, err := strconv.Atoi(e.FormValue("pi"))
	if err!=nil{
		return e.JSON(utils.ErrIpt("分页索引错误",err.Error()))//ErrIpt数据
	}
	if pi<1{
		return e.JSON(utils.ErrIpt("分页索引范围错误",err.Error()))
	}
	ps, err := strconv.Atoi(e.FormValue("ps"))
	if err!=nil{
		return e.JSON(utils.ErrIpt("分页大小错误",err.Error()))//ErrIpt数据
	}
	if ps<1||ps>50{
		ps=6 //超过范围给默认值
	}
	count := model.ParticularCount()
	if count<1{
		return e.JSON(utils.ErrOpt("未查询到数据"))
	}
	mods, err := model.ParticularPage(pi,ps)
	if err!=nil{
		return e.JSON(utils.ErrOpt("未查询到数据",err.Error()))
	}
	return e.JSON(utils.Page("Particular，分页数据",mods,count))
}



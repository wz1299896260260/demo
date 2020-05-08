package Particular

import (

	"github.com/labstack/echo"
	"github.com/zxysilent/utils"
	"web/model"
)

func ParticularAdd(e echo.Context)error{
	ipt := &model.Particular{}
	err := e.Bind(ipt)
	if err!=nil{
		return e.JSON(utils.ErrIpt("输入数据有误",err.Error()))
	}
	//if ipt.Name==""{
	//	return e.JSON(utils.ErrIpt("分类名称不能为空"))
	//}
	if model.ParticularExists(ipt.Uid){
		return e.JSON(utils.ErrIpt("请勿重复添加"))
	}

	err = model.ParticularAdd(ipt)

	if err!=nil{
		return e.JSON(utils.Fail("添加失败",err.Error()))
	}
	return e.JSON(utils.Succ("添加成功"))

}


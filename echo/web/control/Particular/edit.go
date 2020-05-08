package Particular

import (
	"github.com/labstack/echo"
	"github.com/zxysilent/utils"
	"strconv"

	"web/model"
)

//或许某一条数据+修改数据库何在一起等于修改
func ParticularGet(e echo.Context) error {
	//uid := e.Param("uid") //获取path参数
	uid, err := strconv.ParseInt(e.Param("uid"), 10, 64)
	//if err != nil {
	//	return e.JSON(utils.ErrIpt("输入数据有误", err.Error()))
	//}

	mod, err := model.ParticularGet(uid)
	if err != nil {
		return e.JSON(utils.ErrOpt("未查询到数据", err.Error())) //ErrOpt没有数据返回
	}
	return e.JSON(utils.Succ("Particular数据", mod))
}

func ParticularGet4(e echo.Context) error {
	//uid := e.Param("uid") //获取path参数
	name := e.Param("name")
	//if err != nil {
	//	return e.JSON(utils.ErrIpt("输入数据有误", err.Error()))
	//}

	mod, err := model.ParticularGet4(name)
	if err != nil {
		return e.JSON(utils.ErrOpt("未查询到数据", err.Error())) //ErrOpt没有数据返回
	}
	return e.JSON(utils.Succ("Particular数据", mod))
}

//修改数据库
func ParticularEdit(e echo.Context) error {
	ipt := &model.Particular{}
	err := e.Bind(ipt)
	if err != nil {
		return e.JSON(utils.ErrIpt("输入数据有误", err.Error()))
	}
	//if ipt.Name==""{
	//	return e.JSON(utils.ErrIpt("分类名称不能为空"))
	//}

		err = model.ParticularEdit(ipt)
		if err != nil {
			return e.JSON(utils.Fail("修改失败", err.Error()))
		}
		return e.JSON(utils.Succ("修改成功"))




}

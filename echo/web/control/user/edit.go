package user

import (
	"github.com/labstack/echo"
	"github.com/zxysilent/utils"
	"time"
	"web/model"
)

//修改
func UserEdit(e echo.Context)error{
		ipt := &model.User{}
		err := e.Bind(ipt)
		if err != nil {
		return e.JSON(utils.ErrIpt("输入数据有误", err.Error()))
	}
	//	if ipt.Num == "" {
	//	return e.JSON(utils.ErrIpt("账号不能为空"))
	//}
		//if ipt.Name==""{
		//	return e.JSON(utils.ErrIpt("用户名不能为空"))
		//}
		//if ipt.Pass==""{
		//	return e.JSON(utils.ErrIpt("用户密码不能为空"))
		//}
		//if ipt.Email!=""{
		//	return e.JSON(utils.ErrIpt("用户邮箱不能为空"))
		//}
	//	if !IsNum(ipt.Num){
	//	return e.JSON(utils.ErrIpt("账号格式错误"))
	//}
	//	if model.UserExists(ipt.Num) {
	//	return e.JSON(utils.ErrIpt("账号重复"))
	//}
	//	if model.UserExists1(ipt.Name) {
	//	return e.JSON(utils.ErrIpt("昵称重复"))
	//}
		if model.UserExists2(ipt.Password) {
			return e.JSON(utils.ErrIpt("与近期密码相同"))
		}
		ipt.Ctime = time.Now()
		ipt.Password = GetMd5String1(ipt.Password)
	//	if model.UserExists2(ipt.Pass) {
	//	return e.JSON(utils.ErrIpt("与近期密码相同"))
	//}
	//	if !IsEmail(ipt.Email){
	//	return e.JSON(utils.ErrIpt("邮箱格式错误"))
	//}

		err = model.UserEdit(ipt)
		if err != nil {
		return e.JSON(utils.Fail("密码不能与近期相同", err.Error()))
	}
		return e.JSON(utils.Succ("修改成功"))
}

//查询所有
//func UserALL(e echo.Context)error{
//	mods, err := model.UserAll()
//	if err!=nil{
//		return e.JSON(utils.Fail("未查询到数据",err.Error()))
//	}
//
//	return e.JSON(utils.Succ("Particular数据",mods))
//}

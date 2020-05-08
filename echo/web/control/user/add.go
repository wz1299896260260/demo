package user

import (
	"crypto/md5"
	"fmt"
	"github.com/labstack/echo"
	"github.com/zxysilent/utils"
	"io"
	"regexp"
	"time"
	"web/model"
)

//添加，判断是否重复
func UserAdd(e echo.Context) error {
	ipt := &model.User{}
	err := e.Bind(ipt)
	if err != nil {
		return e.JSON(utils.ErrIpt("输入数据有误", err.Error()))
	}
	if ipt.Username == "" {
		return e.JSON(utils.ErrIpt("账号不能为空"))
	}
	//if ipt.Name==""{
	//	return e.JSON(utils.ErrIpt("用户名不能为空"))
	//}
	//if ipt.Pass==""{
	//	return e.JSON(utils.ErrIpt("用户密码不能为空"))
	//}
	//if ipt.Email!=""{
	//	return e.JSON(utils.ErrIpt("用户邮箱不能为空"))
	//}
	if !IsNum(ipt.Username){
		return e.JSON(utils.ErrIpt("账号格式错误"))
	}
	if model.UserExists(ipt.Username) {
		return e.JSON(utils.ErrIpt("账号重复"))
	}
	if model.UserExists1(ipt.Name) {
		return e.JSON(utils.ErrIpt("昵称重复"))
	}
	ipt.Ctime = time.Now()
	ipt.Password = GetMd5String1(ipt.Password)
	//if model.UserExists1(ipt.Pass) {
	//	return e.JSON(utils.ErrIpt("与近期密码相同"))
	//}
	if !IsEmail(ipt.Email){
		return e.JSON(utils.ErrIpt("邮箱格式错误"))
	}

	err = model.UserAdd(ipt)
	if err != nil {
		return e.JSON(utils.Fail("注册失败", err.Error()))
	}
	return e.JSON(utils.Succ("注册成功"))
}



//验证md5，正则表达式
func GetMd5String1(str string) string {
	m := md5.New()
	io.WriteString(m, str)
	arr := fmt.Sprintf("%x", m.Sum(nil))
	return arr
}

func IsEmail(email string) bool {
	if email != "" {
		if isOk, _ := regexp.MatchString("^[_a-z0-9-]+(\\.[_a-z0-9-]+)*@[a-z0-9-]+(\\.[a-z0-9-]+)*(\\.[a-z]{2,4})$", email); isOk {
			return true
		}
	}
	return false
}

func IsNum(num string) bool {
	if num != "" {
		if isOk, _ := regexp.MatchString("[a-zA-Z0-9_]{4,15}$", num); isOk {
			return true
		}
	}
	return false
}

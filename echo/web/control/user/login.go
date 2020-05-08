package user

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/zxysilent/utils"
	"time"

	"web/model"
)

type login struct {
	Username string `json:"username"`
	Password     string `json:"password"`
}

//用户登录
func UserLogin(e echo.Context) error {
	ipt := login{}
	err := e.Bind(&ipt)
	if err != nil {
		return e.JSON(utils.ErrIpt("输入有误", err.Error()))
	}
	//ipt.Pass=GetMd5String1(ipt.Pass)
	mod, err := model.UserLogin(ipt.Username)
	if err != nil {
		return e.JSON(utils.ErrIpt("账号或密码有误", err.Error()))
	}
	if mod.Password != ipt.Password {
		return e.JSON(utils.ErrIpt("账号或密码有误"))
	}

	claims := model.UserClaims{
		Id:   mod.Id,
		Username:  mod.Username,
		Name: mod.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(2 * time.Hour).Unix(), //在。。时候过期
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte("123")) //需要加密的秘钥
	//fmt.Printf("%v %v", ss, err)
	return e.JSON(utils.Succ("登陆成功", ss))

}

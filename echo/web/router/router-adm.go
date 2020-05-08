package router

import (
	"github.com/labstack/echo"
	"web/control/Article"
	"web/control/Particular"
	"web/control/user"
)

//以adm开始必须登录之后才能访问,针对数据，除查询外
func AdmRouter(adm *echo.Group){
	//school
	adm.POST("/particular/add", Particular.ParticularAdd)    //添加

	adm.GET("/particular/del/:id", Particular.ParticularDel) //删除//path参数

	adm.POST("/particular/edit", Particular.ParticularEdit) //修改


	adm.GET("/particular/get/:uid", Particular.ParticularGet)

	adm.GET("/particular/get4/:name", Particular.ParticularGet4)


	adm.POST("/user/edit",user.UserEdit)//用户修改

	//留言
	adm.POST("/article/add",Article.ArticleAdd)
}

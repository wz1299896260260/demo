package router

import (
	"github.com/labstack/echo"
	"web/control/Article"
	"web/control/Particular"
	"web/control/info"
	"web/control/user"
)
//大家都能访问，查询放到这里
//显示
func Api(api *echo.Group){
	api.POST("/login", user.UserLogin)//登录
	//class
	api.GET("/particular/all", Particular.ParticularALL)     //查询信息
	api.GET("/article/all", Article.ArticleAll)				//查询留言

	api.GET("/particular/get/:school", Particular.ParticularSchoolSearch)
	api.GET("/particular/get2/:class", Particular.ParticularClassSearch)
	api.GET("/particular/get3/:name", Particular.ParticularGet) //先在页面获取id，然后跳转到修改页面
	api.GET("/particular/get4/:name", Particular.ParticularGet4)

	api.GET("/particular/page", Particular.ParticularPage)   //分页设置
	//user
	api.POST("/user/add",user.UserAdd)



	//查询
	api.GET("/user/search/:username",user.UserGet)
	api.POST("/user/edit",user.UserEdit)

	//图片
	api.POST("/upload",info.Upload)

}

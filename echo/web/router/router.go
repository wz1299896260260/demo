package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"web/control"
)

var debug = true //debug 模式

func Run() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
		//AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,echo.HeaderAuthorization},//,echo.HeaderAuthorization
	}))

	e.Renderer = renderer
	e.HideBanner = true
	//注册静态目录
	e.Static("/static", "static")
	e.Static("/view", "view")
	e.GET("/", control.Index)
	e.GET("/login.html",control.LoginView)
	e.GET("/upload.html",control.UploadView)
	//分组，访问index.html必须带adm前缀。可以放中间件
	//在这个组里的东西都要先访问中间件
	/*adm:=e.Group("/adm",ServerHeader)//限制用token
	adm.GET("/index.html",control.AdmIndexView)//界面用vue*/


	api:=e.Group("/api")
	Api(api)
	adm:=e.Group("/adm",ServerHeader)//,ServerHeader
	AdmRouter(adm)

	e.Start(":8081")
}

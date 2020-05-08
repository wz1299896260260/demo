package control

import "github.com/labstack/echo"

func Index(e echo.Context)error{
	return e.Redirect(302,"/upload.html")//重定向
}

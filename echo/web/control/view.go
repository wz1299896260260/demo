package control

import "github.com/labstack/echo"


//login界面
func LoginView(e echo.Context)error{
	return e.Render(200,"login.html",nil)
}

//Index界面
func AdmIndexView(e echo.Context)error{
	return e.Render(200,"index.html",nil)
}

//Upload界面
func UploadView(e echo.Context)error{
	return e.Render(200,"upload.html",nil)
}

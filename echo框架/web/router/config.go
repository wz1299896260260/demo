package router

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/zxysilent/utils"
	"web/model"

	"io"
	"text/template"

)

//加载到内存
// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	//如果debug模式则重新加载
	if debug {
		t.templates = template.Must(template.ParseFiles("./view/login.html", "./view/index.html","./view/upload.html"))
	}
	return t.templates.ExecuteTemplate(w, name, data)
}

var renderer = &TemplateRenderer{
	templates: template.Must(template.ParseFiles("./view/login.html", "./view/index.html","./view/upload.html")),
}


//中间件
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
	//e.Response().Header().Set(echo.HeaderServer,"Echo/999")
		//tokenString := e.FormValue("token")
		tokenString := e.Request().Header.Get("token")
		claims := model.UserClaims{}
		token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("123"), nil
		})
		if err == nil && token.Valid {
			//验证通过
			e.Set("uid",claims.Id)
			return next(e)
		} else {
			return e.JSON(utils.ErrJwt("jwt token验证失败"))
		}
	}
}



//func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
//	return func(ctx echo.Context) error {
//		// query form 查找 token
//		tokenString := ctx.FormValue("token")
//		if tokenString == "" {
//			// header 查找token
//			tokenString = ctx.Request().Header.Get(echo.HeaderAuthorization)
//			if tokenString == "" {
//				ctx.JSON(utils.ErrJwt(`请重新登陆`, `未发现jwt`))
//				return nil
//			}
//			// Bearer token
//			tokenString = tokenString[7:] //len("Bearer ")
//		}
//		jwtAuth := &model.UserClaims{}
//		jwt, err := jwt.ParseWithClaims(tokenString, jwtAuth, func(token *jwt.Token) (interface{}, error) {
//			return []byte("zxy.sil.ent"), nil
//		})
//		if err == nil && jwt.Valid {
//			ctx.Set("auth", jwtAuth)
//			ctx.Set("uid", jwtAuth.Id)
//		} else {
//			return ctx.JSON(utils.ErrJwt(`请重新登陆","jwt验证失败`))
//		}
//		// 自定义头
//		ctx.Response().Header().Set(echo.HeaderServer, "zxysilent")
//		return next(ctx)
//	}
//}
//func MidJwt(next echo.HandlerFunc) echo.HandlerFunc {
//	return func(ctx echo.Context) error {
//		tokenRaw := ctx.FormValue("token") // query/form 查找 token
//		if tokenRaw == "" {
//			tokenRaw = ctx.Request().Header.Get(echo.HeaderAuthorization) // header 查找token
//			if tokenRaw == "" {
//				ctx.JSON(utils.ErrJwt(`请重新登陆`, `未发现jwt`))
//				return nil
//			}
//			//tokenRaw = tokenRaw[7:] // Bearer token len("Bearer ")==7
//		}
//		jwtAuth, err := Verify(tokenRaw, App.Jwtkey)
//		if err == nil {
//			ctx.Set("auth", jwtAuth)
//			ctx.Set("uid", jwtAuth.Id)
//		} else {
//			return ctx.JSON(utils.ErrJwt(`请重新登陆","jwt验证失败`))
//		}
//		// 自定义头
//		return next(ctx)
//	}
//}



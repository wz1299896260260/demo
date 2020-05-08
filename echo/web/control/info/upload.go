package info

import (

	"github.com/labstack/echo"
	"github.com/zxysilent/utils"
	"io"
	"path"
	"web/model"

	"os"
	"time"



)

func Upload(c echo.Context) error {
	note:=c.FormValue("note")
	// Source 从表单获取数据流
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()//打开文件
	if err != nil {
		return err
	}
	defer src.Close()

	////Destination
	//dst, err := os.Create(file.Filename)
	//if err != nil {
	//	return err
	//}
	//defer dst.Close()
	//
	////Copy
	//_, err = io.Copy(dst, src)
	//if err != nil {
	//	return err
	//}
	os.Mkdir("./static/",0666)
	now := time.Now()
	name := now.Format("2006-01-02150405") + path.Ext(file.Filename)
	out, err := os.Create("./static/" + name) //创建本地目录
	if err!=nil{
		return c.JSON(utils.Fail("添加失败"))
	}
	io.Copy(out, src)  //将上传的数据拷贝到本地文件中
	out.Close()

	mod := model.Info{
		Name: file.Filename,
		Path: "static/"+name,
		Note: note,
		Unix: time.Now().Unix(),
	}
	err = model.InfoAdd(&mod)
	if err!=nil{
		return c.JSON(utils.Fail("添加失败1"))
	}
	return c.JSON(utils.Succ("添加成功"))
}
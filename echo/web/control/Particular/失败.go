package Particular



/*//或许某一条数据+修改数据库何在一起等于修改
func ClassGet(e echo.Context)error{
	//e.Param("id") //获取path参数
	id, err := strconv.ParseInt(e.Param("id"), 10, 64)
	if err!=nil {
		return e.JSON(utils.ErrIpt("输入数据有误", err.Error()))
	}
	mod, err := model.ClassGet(id)
	if err!=nil {
		return e.JSON(utils.ErrOpt("未查询到数据", err.Error()))//ErrOpt没有数据返回
	}
	return e.JSON(utils.Succ("分类数据",mod))
}
//修改数据库
func ClassEdit(e echo.Context)error{
	ipt := &model.Class{}
	err := e.Bind(ipt)
	if err!=nil{
		return e.JSON(utils.ErrIpt("输入数据有误",err.Error()))
	}
	if ipt.Name==""{
		return e.JSON(utils.ErrIpt("分类名称不能为空"))
	}
	err = model.ClassEdit(ipt)
	if err!=nil{
		return e.JSON(utils.Fail("修改失败",err.Error()))
	}
	return e.JSON(utils.Succ("修改成功"))

}*/


/*//分页和数据总数一起是查询所有分类
func ClassALL(e echo.Context)error{
	mods, err := model.ClassAll()
	if err!=nil{
		return e.JSON(utils.Fail("未查询到数据",err.Error()))
	}
	return e.JSON(utils.Succ("分类数据",mods))
}

//分页查询
func ClassPage(e echo.Context)error{
	pi, err := strconv.Atoi(e.FormValue("pi"))
	if err!=nil{
		return e.JSON(utils.ErrIpt("分页索引错误",err.Error()))//ErrIpt数据
	}
	if pi<1{
		return e.JSON(utils.ErrIpt("分页索引范围错误",err.Error()))
	}
	ps, err := strconv.Atoi(e.FormValue("ps"))
	if err!=nil{
		return e.JSON(utils.ErrIpt("分页大小错误",err.Error()))//ErrIpt数据
	}
	if ps<1||ps>50{
		 ps=6 //超过范围给默认值
	}
	count := model.ClassCount()
	if count<1{
		return e.JSON(utils.ErrOpt("未查询到数据"))
	}
	mods, err := model.ClassPage(pi,ps)
	if err!=nil{
		return e.JSON(utils.ErrOpt("未查询到数据",err.Error()))
	}
	return e.JSON(utils.Page("分类，分页数据",mods,count))
}*/

//添加
/*func ClassAdd(e echo.Context)error{
	ipt := &model.Class{}
	err := e.Bind(ipt)
	if err!=nil{
		return e.JSON(utils.ErrIpt("输入数据有误",err.Error()))
	}
	if ipt.Name==""{
		return e.JSON(utils.ErrIpt("分类名称不能为空"))
	}
	err = model.ClassAdd(ipt)
	if err!=nil{
		return e.JSON(utils.Fail("添加失败",err.Error()))
	}
		return e.JSON(utils.Succ("添加成功"))

}*/

//删除分类
/*func ClassDel(e echo.Context)error{
	id, err := strconv.ParseInt(e.Param("id"), 10, 64)
	if err!=nil {
		return e.JSON(utils.ErrIpt("输入数据有误", err.Error()))
	}
	err = model.ClassDel(id)
	if err!=nil {
		return e.JSON(utils.Fail("删除失败", err.Error()))
	}
	return e.JSON(utils.Succ("删除成功"))
}*/


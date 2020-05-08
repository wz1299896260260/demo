package model

import "errors"

//info 一个照片信息
type Info struct {
	Id   int64
	Name string
	Path string //保存路径
	Note string
	Unix int64 //时间
}

//添加
func InfoAdd(mod *Info) error {
	result, err := DB.Exec("insert into Info(`name`,path,unix,note) values (?,?,?,?)",
		mod.Name, mod.Path,mod.Unix,mod.Note)
	if err != nil {
		return err
	}
	id, _ := result.LastInsertId()
	if id < 1 {
		return errors.New("添加失败")
	}
	return nil
}

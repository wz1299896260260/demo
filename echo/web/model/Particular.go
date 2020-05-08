package model

import (
	"errors"

)

//数据表
//Particular类别
type Particular struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	City   string `json:"city"`
	School string `json:"school"`
	Class  string `json:"class"`
	Age    string `json:"age"`
	Sex    string `json:"sex"`
	QQ     string `json:"qq"`
	Uid    int64  `json:"uid"`
}

//分页
func ParticularPage(pi, ps int) ([]Particular, error) {
	mods := make([]Particular, 0, ps)
	err := DB.Select(&mods, "select * from particular limit ?,?", (pi-1)*ps, ps)
	return mods, err
}

//需要分页总数才能计算量

//数据总数
func ParticularCount() int {
	count := 0
	DB.Get(&count, "select count(id) as count from particular")
	return count
}

//查询所有
func ParticularAll() ([]Particular, error) {
	mods := make([]Particular, 0, 8)
	err := DB.Unsafe().Select(&mods, "select distinct school  from particular ")
	return mods, err
}

//查询学校
func ParticularSchoolSearchGet(school string) ([]Particular, error) {
	//mod := Particular{}
	//distinct class 只允许class不同数据显示
	mods := make([]Particular, 0, 8)
	err := DB.Unsafe().Select(&mods, "select distinct class from particular where school=? ", school)
	return mods, err
}

//查询班级
func ParticularClassSearchGet(class string) ([]Particular, error) {
	//mod := Particular{}
	//distinct class 只允许class不同数据显示
	mods := make([]Particular, 0, 8)
	err := DB.Unsafe().Select(&mods, "select * from particular where class=?", class)
	return mods, err
}

//查询某一条//编辑要用这个查询id
func ParticularGet(uid int64) (*Particular, error) {
	mod := &Particular{}
	err := DB.Get(mod, "select * from particular where uid=? limit 1",uid)
	return mod, err
}

func ParticularGet4(name string) (*Particular, error) {
	mod := &Particular{}
	err := DB.Get(mod, "select * from particular where name=? limit 1",name)
	return mod, err
}

//添加
func ParticularAdd(mod *Particular) error {
	tx, _ := DB.Unsafe().Begin() //开启一个事务，保险
	result, err := tx.Exec("insert into particular(`name`,`phone`,`city`,`school`,`class`,`age`,`sex`,`qq`,`uid`)values(?,?,?,?,?,?,?,?,?)",
		mod.Name, mod.Phone, mod.City, mod.School, mod.Class, mod.Age, mod.Sex, mod.QQ, mod.Uid)
	if err != nil {
		//回滚,撤销
		tx.Rollback()
		return err
	}
	//update, insert, or delete都会返回受影响的行数
	row, _ := result.RowsAffected()
	if row < 1 {
		//未插入  回滚
		tx.Rollback()
		return errors.New("rows affected<1")
	}
	//手动提交，写入数据  commit之后不可回滚
	tx.Commit()
	return nil
}

//删除
func ParticularDel(id int64) error {
	tx, _ := DB.Begin() //开启一个事务，保险，
	result, err := tx.Exec("delete from particular where id=?", id)
	if err != nil {
		//回滚,撤销
		tx.Rollback()
		return err
	}
	//update, insert, or delete都会返回受影响的行数
	row, _ := result.RowsAffected()
	if row < 1 {
		//回滚
		tx.Rollback()
		return errors.New("rows affected<1")
	}
	//手动提交，写入数据
	tx.Commit()
	return nil
}

//修改
func ParticularEdit(mod *Particular) error {
	tx, _ := DB.Begin() //开启一个事务，保险
	result, err := tx.Exec("update particular set `name`=?,`phone`=?,`city`=?,`school`=?,`class`=?,age=?,sex=?,qq=? where `id`=?",
		mod.Name, mod.Phone, mod.City, mod.School, mod.Class, mod.Age, mod.Sex, mod.QQ, mod.ID)
	if err != nil {
		//回滚,撤销
		tx.Rollback()
		return err
	}
	//update, insert, or delete都会返回受影响的行数
	row, _ := result.RowsAffected()
	if row < 1 {
		//回滚
		tx.Rollback()
		return errors.New("rows affected<1")
	}
	//提交，写入数据
	tx.Commit()
	return nil
}

func ParticularExists(uid int64) bool {
	mod := Particular{}
	err := DB.Get(&mod, "select * from particular where uid=? ", uid)
	if err != nil {
		return false
	}
	return true
}
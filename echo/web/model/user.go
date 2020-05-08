package model

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

//user数据表 表
type User struct {
	Id       int       `json:"id"`
	Username string    `json:"username"`
	Name     string    `json:"name"`
	Password string    `json:"password"`
	Phone    string    `json:"phone"`
	Email    string    `json:"email"`
	Ctime    time.Time `json:"ctime"`
}

//token 携带的数据
type UserClaims struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	jwt.StandardClaims
}

//判断用户是否存在
func UserExists(username string) bool {
	mod := User{}
	err := DB.Get(&mod, "select * from user where username=? ", username)
	if err != nil {
		return false
	}
	return true
}

func UserExists1(name string) bool {
	mod := User{}
	err := DB.Get(&mod, "select * from user where name=? ", name)
	if err != nil {
		return false
	}
	return true
}
func UserExists2(password string) bool {
	mod := User{}
	err := DB.Get(&mod, "select * from user where password=? ", password)
	if err != nil {
		return false
	}
	return true
}

//用户登录，查询
func UserLogin(username string) (User, error) {
	mod := User{}
	err := DB.Get(&mod, "select * from user where username=? limit 1", username)
	return mod, err
}

//用户添加
func UserAdd(mod *User) error {
	tx, _ := DB.Begin() //开启一个事务，保险

	result, err := tx.Exec("insert into user(username,`name`,password,phone,email,ctime)values(?,?,?,?,?,?)",
		mod.Username, mod.Name, mod.Password, mod.Phone, mod.Email, mod.Ctime)
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

//修改等于查询id等到内容+修改数据库
func UserGet(username string) (User, error) {
	mod := User{}
	err := DB.Unsafe().Get(&mod, "select * from user where username=? ", username)
	return mod, err
}
func AdminUserGet(id int64) (User, error) {
	mod := User{}
	err := DB.Unsafe().Get(&mod, "select * from user where id='5' ", id)
	return mod, err
}

func UserEdit(mod *User) error {
	tx, _ := DB.Begin() //开启一个事务，保险
	result, err := tx.Exec("update user set `name`=?,password=?,phone=?,email=? where username=?", mod.Name, mod.Password, mod.Phone, mod.Email, mod.Username)
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

//func UserAll() ([]User, error) {
//	mods := make([]User, 0, 8)
//	err := DB.Unsafe().Select(&mods, "select * from user ")
//	return mods, err
//}

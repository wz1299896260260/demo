package model

import (
	"errors"
)

type Article struct {
	//表的内容
	Id      int64  `json:"id"`
	Content string `json:"content"`
	Uid     int64  `json:"uid"`
	User
}



//count 数据总数
func ArticleCount() int {
	count := 0
	DB.Get(&count, "select count(id) as count from class")
	return count
}

//分页数据

func ArticlePage(pi, ps int) ([]Article, error) {
	mods := make([]Article, 0, ps)
	err := DB.Select(&mods, "select * from article limit ?,?", (pi-1)*ps, ps)
	return mods, err
}

//添加信息
func ArticleAdd(mod *Article) error {
	tx, _ := DB.Unsafe().Begin() //开启一个事务，保险
	result, err := tx.Exec("insert into article(`content`,`uid`)values(?,?)",
		mod.Content, mod.Uid)
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

//查询所有
func ArticleAll() ([]Article, error) {
	mods := make([]Article, 0, 8)
	//err := DB.Unsafe().Select(&mods, "select * from article ")
	err := DB.Unsafe().Select(&mods, "select * from user,article where user.id=article.uid ")
	//err := DB.Unsafe().Select(&mods,"select name from user where id in(select uid from article) union select content from article where uid in(select id from user)")
	return mods, err
}

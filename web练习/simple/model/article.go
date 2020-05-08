package model

import (
	"time"
)

//表单
type Article struct {
	Id      int64     `json:"id"`
	Title   string    `json:"title"`
	Author  string    `json:"author"`
	Content string    `json:"content"`
	Hits    int       `json:"hits"`
	Utime   time.Time `json:"utime"`
}

//查询一条，数据库操作
func ArticleGet(id int64) (Article, error) {
	mod := Article{}
	err := DB.Unsafe().Get(&mod, "select * from Article where id=? limit 1 ", id)
	return mod, err
}

//查询集合
func ArticleList() ([]Article, error) {
	mods := make([]Article, 0, 10)
	err := DB.Unsafe().Select(&mods, "select * from Article order by id desc limit 10 ")
	return mods, err
}

//del 删除
func ArticleDel(id int64) bool {
	res, _ := DB.Exec("delete from Article where id=?", id)
	if res == nil {
		return false
	}
	rows, _ := res.RowsAffected() //影响几行
	if rows >= 1 {
		return true
	}
	return false
}

//insert 插入
func ArticleAdd(mod *Article) error {
	_, err := DB.Exec("insert into Article (title,author,content,hits,utime)values(?,?,?,?,?)", mod.Title, mod.Author, mod.Content, mod.Hits, mod.Utime)
	return err
}

//edit
func ArticleEdit(mod *Article)error{
	_, err := DB.Exec("update Article set title=?,author=?,content=?,hits=? where id=?", mod.Title, mod.Author, mod.Content, mod.Hits, mod.Id)
	return err
}

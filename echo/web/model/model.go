package model

import (
	"github.com/jmoiron/sqlx"
	"log"
	_"github.com/go-sql-driver/mysql"
)

var DB *sqlx.DB

func init(){
	//连接数据库
	db, err := sqlx.Open(`mysql`, `root:123456@tcp(localhost:3306)/news?charset=utf8&parseTime=true`)
	if err!=nil{
		log.Fatalln(err.Error())
	}
	//测试数据库正常工作
	if err=db.Ping();err!=nil{
		log.Fatalln(err.Error())
	}
	DB=db
	
}

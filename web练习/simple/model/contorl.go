package model

import (
	"github.com/jmoiron/sqlx"
	_"github.com/go-sql-driver/mysql"
	"log"
)

//链接数据库

var DB *sqlx.DB

func init(){
	db, err := sqlx.Open(`mysql`, `root:123456@tcp(localhost:3306)/news?charset=utf8&parseTime=true`)
	if err!=nil{
		log.Fatalln(err.Error())
	}
	DB=db

}

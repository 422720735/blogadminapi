package dbops

import (
	"blogadminapi/lib"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DbConn *sql.DB
	err    error
)

func Init() {
	username := lib.Conf.Read("mysql", "username")
	password := lib.Conf.Read("mysql", "password")
	dataname := lib.Conf.Read("mysql", "dataname")
	port := lib.Conf.Read("mysql", "port")
	host := lib.Conf.Read("mysql", "host")
	dns := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dataname + "?parseTime=true"
	DbConn, err = sql.Open("mysql", dns)
	//dbConn, err =sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/gin_book?charset=utf8")
	if err != nil {
		panic(err)
	}
	fmt.Print("连接成功")
	DbConn.SetConnMaxLifetime(20)
	DbConn.SetMaxIdleConns(20)
}

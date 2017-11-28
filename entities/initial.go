package entities

import (
	//"database/sql"
	"github.com/go-xorm/xorm"

	_ "github.com/go-sql-driver/mysql"
)

// 一个Orm引擎称为Engine，一个Engine一般只对应一个数据库
var engine *xorm.Engine

func init() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8")
  checkErr(err)

  // 同步能够部分智能的根据结构体的变动检测表结构的变动，并自动同步
  err2 := engine.Sync2(new(UserInfo))
  checkErr(err2)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

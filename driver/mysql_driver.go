package driver

import (
	"github.com/jinzhu/gorm"
	"os"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type MysqlDB struct {
	SQL *gorm.DB
}

var Mysql = &MysqlDB{}

func Connect() (*MysqlDB){
	db, err := gorm.Open("mysql", os.Getenv("DB_CONFIG"))
	if err != nil {
		panic(err)
	}
	Mysql.SQL = db
	return Mysql
}
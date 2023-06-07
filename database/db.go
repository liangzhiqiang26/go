package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func InitDB() error {
	var err error
	DB, err = sqlx.Open("mysql", "cjw:cjw@tcp(172.168.30.212:3308)/scrm_assignment")
	if err != nil {
		return err
	}

	// 配置连接池
	DB.SetMaxOpenConns(10) // 设置最大打开连接数
	DB.SetMaxIdleConns(5)  // 设置最大空闲连接数

	// 测试数据库连接
	err = DB.Ping()
	if err != nil {
		return err
	}

	return nil
}

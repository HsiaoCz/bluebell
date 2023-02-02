package mysql

import (
	"bluebell/setting"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var db *sqlx.DB

func Init() (err error) {
	mcg := setting.Conf.MysqlConfig
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", mcg.User, mcg.Password, mcg.Host, mcg.Port, mcg.DbName)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect DB failed,err:%v\n", zap.Error(err))
		return
	}
	db.SetMaxOpenConns(mcg.MaxOpenConns)
	db.SetMaxIdleConns(mcg.MaxIdleConns)
	return err
}

func Close() {
	_ = db.Close()
}

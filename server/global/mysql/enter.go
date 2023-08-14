package mysql

import (
	"CTFe/server/global/config"
	"CTFe/server/util/log"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	info := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		config.GlobalConfig.MysqlConfig.Username,
		config.GlobalConfig.MysqlConfig.Password,
		config.GlobalConfig.MysqlConfig.Host,
		config.GlobalConfig.MysqlConfig.Port,
		config.GlobalConfig.MysqlConfig.Database,
	)

	var err error
	db, err = sql.Open("mysql", info)
	if err != nil {
		log.InfoLogger.Println(err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.InfoLogger.Println(err.Error())
	}
}

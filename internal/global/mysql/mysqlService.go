package mysql

import (
	"CTFe/internal/global/config"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os/user"
)

var (
	err error
	db  *sql.DB
)

func init() {
	info := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		config.GlobalConfig.MySqlConfig.Username,
		config.GlobalConfig.MySqlConfig.Password,
		config.GlobalConfig.MySqlConfig.Host,
		config.GlobalConfig.MySqlConfig.Port,
		config.GlobalConfig.MySqlConfig.Database,
	)
	db, err = sql.Open("mysql", info)
	if err != nil {
		fmt.Println(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
}

func InsertUser(user user.User) {

}

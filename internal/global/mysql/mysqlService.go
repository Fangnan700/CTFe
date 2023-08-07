package mysql

import (
	"CTFe/internal/global/config"
	"CTFe/internal/models"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func init() {
	var err error

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

/*
	用户表相关操作
*/

// InsertUser 新增用户
func InsertUser(user models.Users) error {
	var (
		err  error
		tx   *sql.Tx
		stmt *sql.Stmt
	)

	// 开启事务
	tx, err = db.Begin()
	if err != nil {
		return err
	}

	//准备sql语句
	stmt, err = tx.Prepare("INSERT INTO users (user_id, user_name, user_pwd, user_sex, email, phone, school, student_num, create_time) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);")
	if err != nil {
		return err
	}

	//设置参数以及执行sql语句
	_, err = stmt.Exec(user.UserId, user.UserName, user.UserPwd, user.UserSex, user.Email, user.Phone, user.School, user.StudentNum, user.CreateTime)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

// SelectLastUserId 查询最后一个用户ID
func SelectLastUserId() (int64, error) {
	var (
		err        error
		row        *sql.Row
		lastUserId int64
	)

	row = db.QueryRow("SELECT user_id FROM users ORDER BY user_id DESC LIMIT 1;")
	err = row.Scan(&lastUserId)
	if err != nil {
		return -1, err
	}
	return lastUserId, nil
}

// SelectUsers 查询用户
func SelectUsers(keyword interface{}) ([]models.Users, error) {
	var (
		err   error
		rows  *sql.Rows
		users []models.Users
	)

	rows, err = db.Query("SELECT * FROM users WHERE user_id LIKE ? OR user_name LIKE ? OR email LIKE ? OR phone LIKE ? OR school LIKE ? OR student_num LIKE ?", keyword, keyword, keyword, keyword, keyword, keyword)
	for rows.Next() {
		var user models.Users
		err = rows.Scan(&user.UserId, &user.UserName, &user.UserPwd, &user.UserSex, &user.Email, &user.Phone, &user.School, &user.StudentNum, &user.CreateTime)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}

	return users, nil
}

// SelectAllUsers 查询所有用户
func SelectAllUsers() ([]models.Users, error) {
	var (
		err   error
		rows  *sql.Rows
		users []models.Users
	)

	rows, err = db.Query("SELECT * FROM users ORDER BY user_id LIMIT 1, 65536;")
	for rows.Next() {
		var user models.Users
		err = rows.Scan(&user.UserId, &user.UserName, &user.UserPwd, &user.UserSex, &user.Email, &user.Phone, &user.School, &user.StudentNum, &user.CreateTime)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}

	return users, nil
}

// SelectUserById 根据用户id查询
func SelectUserById(userId interface{}) (models.Users, error) {
	var (
		err  error
		row  *sql.Row
		user models.Users
	)

	row = db.QueryRow("SELECT * FROM users WHERE user_id = ?;", userId)
	err = row.Scan(&user.UserId, &user.UserName, &user.UserPwd, &user.UserSex, &user.Email, &user.Phone, &user.School, &user.StudentNum, &user.CreateTime)
	if err != nil {
		return user, err
	}

	return user, nil
}

// SelectUserByEmailOrPhone 根据用户邮箱/手机查询
func SelectUserByEmailOrPhone(email interface{}, phone interface{}) (models.Users, error) {
	var (
		err  error
		row  *sql.Row
		user models.Users
	)

	row = db.QueryRow("SELECT * FROM users WHERE email = ? OR phone = ?;", email, phone)
	err = row.Scan(&user.UserId, &user.UserName, &user.UserPwd, &user.UserSex, &user.Email, &user.Phone, &user.School, &user.StudentNum, &user.CreateTime)
	if err != nil {
		return user, err
	}

	return user, nil
}

// UpdateUser 更新用户信息
func UpdateUser(users models.Users) error {
	_, err := db.Exec("UPDATE users SET user_name = ?, user_pwd = ?, user_sex = ?, email = ?, phone = ?, school = ?, student_num = ? WHERE user_id = ?", users.UserName, users.UserPwd, users.UserSex, users.Email, users.Phone, users.School, users.StudentNum, users.UserId)
	return err
}

// DeleteUser 删除用户
func DeleteUser(userId interface{}) error {
	_, err := db.Exec("DELETE FROM users WHERE user_id = ?", userId)
	return err
}

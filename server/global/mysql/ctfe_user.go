package mysql

import (
	"CTFe/server/model/database"
	"database/sql"
)

/*
	@用户模块
*/

// InsertUser 创建用户
func InsertUser(user database.User) error {
	var (
		err  error
		tx   *sql.Tx
		stmt *sql.Stmt
	)

	tx, err = db.Begin()
	if err != nil {
		return err
	}

	stmt, err = tx.Prepare("INSERT INTO `ctfe_user`(uuid, user_id, user_name, user_pwd, user_sex, user_email, user_phone, user_school, create_time) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(user.UUID, user.UserId, user.UserName, user.UserPwd, user.UserSex, user.UserEmail, user.UserPhone, user.UserSchool, user.CreateTime)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

// DeleteUser 删除用户
func DeleteUser(user database.User) error {
	var (
		err  error
		tx   *sql.Tx
		stmt *sql.Stmt
	)

	tx, err = db.Begin()
	if err != nil {
		return err
	}

	stmt, err = tx.Prepare("DELETE FROM ctfe_user WHERE user_id = ?;")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(user.UserId)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

// UpdateUser 更新用户
func UpdateUser(user database.User) error {
	var (
		err  error
		tx   *sql.Tx
		stmt *sql.Stmt
	)

	tx, err = db.Begin()
	if err != nil {
		return err
	}

	stmt, err = tx.Prepare("UPDATE `ctfe_user` SET user_name = ?, user_sex = ?, user_pwd = ?, user_email = ?, user_phone = ?, user_school = ? WHERE user_id = ?;")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(user.UserName, user.UserSex, user.UserPwd, user.UserEmail, user.UserPhone, user.UserSchool, user.UserId)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

// SelectUser 查询用户
func SelectUser(keyword interface{}) ([]database.User, error) {
	var (
		err   error
		rows  *sql.Rows
		users []database.User
	)

	if keyword == nil {
		rows, err = db.Query("SELECT * FROM `ctfe_user` ORDER BY user_id;")
	} else {
		rows, err = db.Query("SELECT * FROM `ctfe_user` WHERE uuid = ? OR user_id = ? OR user_name LIKE ? OR user_email LIKE ? OR user_phone LIKE ? OR user_school LIKE ?;", keyword, keyword, keyword, keyword, keyword, keyword)
	}
	for rows.Next() {
		var u database.User
		err = rows.Scan(&u.UUID, &u.UserId, &u.UserName, &u.UserPwd, &u.UserSex, &u.UserEmail, &u.UserPhone, &u.UserSchool, &u.CreateTime)
		if err != nil {
			return users, err
		}
		users = append(users, u)
	}

	return users, nil
}

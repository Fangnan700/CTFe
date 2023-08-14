package mysql

import (
	"CTFe/server/model/database"
	"database/sql"
)

/*
	@管理员模块
*/

// InsertAdministrator 添加管理员
func InsertAdministrator(admin database.Admin) error {
	var (
		err  error
		tx   *sql.Tx
		stmt *sql.Stmt
	)

	tx, err = db.Begin()
	if err != nil {
		return err
	}

	stmt, err = tx.Prepare("INSERT INTO `ctfe_admin`(admin_id, user_id) VALUES (?, ?);")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(admin.AdminId, admin.UserId)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

// DeleteAdministrator 删除管理员
func DeleteAdministrator(admin database.Admin) error {
	var (
		err  error
		tx   *sql.Tx
		stmt *sql.Stmt
	)

	tx, err = db.Begin()
	if err != nil {
		return err
	}

	stmt, err = tx.Prepare("DELETE FROM ctfe_admin WHERE admin_id = ? AND user_id = ?;")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(admin.AdminId, admin.UserId)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

// SelectAdministrator 查找管理员
func SelectAdministrator(ID interface{}) ([]database.Admin, error) {
	var (
		err    error
		rows   *sql.Rows
		admins []database.Admin
	)

	if ID == nil {
		rows, err = db.Query("SELECT * FROM `ctfe_admin` ORDER BY admin_id;")
	} else {
		rows, err = db.Query("SELECT * FROM `ctfe_admin` WHERE admin_id = ? OR user_id = ?;", ID, ID)
	}
	for rows.Next() {
		var a database.Admin
		err = rows.Scan(&a.AdminId, &a.UserId)
		if err != nil {
			return admins, err
		}
		admins = append(admins, a)
	}

	return admins, nil
}

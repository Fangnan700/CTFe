package mysql

import (
	"CTFe/server/model/database"
	"database/sql"
)

/*
	@团队模块
*/

// CreateGroup 创建团队
func CreateGroup(group database.Group) error {
	var (
		err  error
		tx   *sql.Tx
		stmt *sql.Stmt
	)

	tx, err = db.Begin()
	if err != nil {
		return err
	}

	stmt, err = tx.Prepare("INSERT INTO `ctfe_group`(group_id, group_name, group_intro, competition_id) VALUES (?, ?, ?, ?);")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(group.GroupId, group.GroupName, group.GroupIntro, group.CompetitionId)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

// DeleteGroup 删除团队
func DeleteGroup(group database.Group) error {
	var (
		err  error
		tx   *sql.Tx
		stmt *sql.Stmt
	)

	tx, err = db.Begin()
	if err != nil {
		return err
	}

	stmt, err = tx.Prepare("DELETE FROM ctfe_group WHERE group_id = ?;")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(group.GroupId)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

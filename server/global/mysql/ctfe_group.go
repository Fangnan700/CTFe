package mysql

import (
	"CTFe/server/model/database"
	"database/sql"
	"fmt"
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

// UpdateGroup 更新团队
func UpdateGroup(group database.Group) error {
	var (
		err  error
		tx   *sql.Tx
		stmt *sql.Stmt
	)

	tx, err = db.Begin()
	if err != nil {
		return err
	}

	stmt, err = tx.Prepare("UPDATE `ctfe_group` SET group_name = ?, group_intro = ? WHERE group_id = ?;")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(group.GroupName, group.GroupIntro, group.GroupId)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

// SelectGroup 查询团队
func SelectGroup(keyword interface{}) ([]database.Group, error) {
	var (
		err    error
		rows   *sql.Rows
		groups []database.Group
	)

	if keyword == nil {
		rows, err = db.Query("SELECT * FROM `ctfe_group` ORDER BY group_id;")
	} else {
		rows, err = db.Query("SELECT * FROM `ctfe_group` WHERE group_id = ? OR ctfe_group.group_name LIKE ?;", keyword, keyword)
	}
	for rows.Next() {
		var g database.Group
		err = rows.Scan(&g.CompetitionId, &g.GroupName, &g.GroupIntro, &g.CompetitionId)
		if err != nil {
			return groups, err
		}
		groups = append(groups, g)
	}

	return groups, nil
}

// JoinGroup 加入团队
func JoinGroup(participation database.Participation) error {
	var (
		err  error
		tx   *sql.Tx
		stmt *sql.Stmt
	)

	tx, err = db.Begin()
	if err != nil {
		return err
	}

	stmt, err = tx.Prepare("INSERT INTO `ctfe_participation`(participation_id, group_id, user_id, competition_id, is_admin) VALUES (?, ?, ?, ?, ?);")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(participation.ParticipationId, participation.GroupId, participation.UserId, participation.CompetitionId, participation.IsAdmin)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

// SelectParticipation 查询团队关系
func SelectParticipation(competitionId interface{}, userId interface{}) ([]database.Participation, error) {
	var (
		err            error
		rows           *sql.Rows
		participations []database.Participation
	)

	fmt.Println(competitionId, userId)

	rows, err = db.Query("SELECT * FROM `ctfe_participation` WHERE competition_id LIKE ? AND user_id LIKE ?;", competitionId, userId)

	fmt.Println(rows)
	for rows.Next() {
		var p database.Participation
		err = rows.Scan(&p.ParticipationId, &p.GroupId, &p.UserId, &p.IsAdmin)
		if err != nil {
			return participations, err
		}
		participations = append(participations, p)
	}

	return participations, nil
}

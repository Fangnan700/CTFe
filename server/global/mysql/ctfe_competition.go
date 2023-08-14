package mysql

import (
	"CTFe/server/model/database"
	"database/sql"
)

/*
	@比赛模块
*/

// InsertCompetition 创建比赛
func InsertCompetition(competition database.Competition) error {
	var (
		err  error
		tx   *sql.Tx
		stmt *sql.Stmt
	)

	tx, err = db.Begin()
	if err != nil {
		return err
	}

	stmt, err = tx.Prepare("INSERT INTO `ctfe_competition`(competition_id, competition_name, description, start_time, left_time) VALUES (?, ?, ?, ?, ?);")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(competition.CompetitionId, competition.CompetitionName, competition.Description, competition.StartTime, competition.LeftTime)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

// DeleteCompetition 删除比赛
func DeleteCompetition(competition database.Competition) error {
	var (
		err  error
		tx   *sql.Tx
		stmt *sql.Stmt
	)

	tx, err = db.Begin()
	if err != nil {
		return err
	}

	stmt, err = tx.Prepare("DELETE FROM ctfe_competition WHERE competition_id = ?;")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(competition.CompetitionId)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

// UpdateCompetition 更新比赛
func UpdateCompetition(competition database.Competition) error {
	var (
		err  error
		tx   *sql.Tx
		stmt *sql.Stmt
	)

	tx, err = db.Begin()
	if err != nil {
		return err
	}

	stmt, err = tx.Prepare("UPDATE `ctfe_competition` SET competition_name = ?, description = ?, start_time = ?, left_time = ? WHERE competition_id = ?;")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(competition.CompetitionName, competition.Description, competition.StartTime, competition.LeftTime, competition.CompetitionId)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

// SelectCompetition 查询比赛
func SelectCompetition(keyword interface{}) ([]database.Competition, error) {
	var (
		err          error
		rows         *sql.Rows
		competitions []database.Competition
	)

	if keyword == nil {
		rows, err = db.Query("SELECT * FROM `ctfe_competition` ORDER BY competition_id;")
	} else {
		rows, err = db.Query("SELECT * FROM `ctfe_competition` WHERE competition_id = ? OR competition_name LIKE ?;", keyword, keyword)
	}
	for rows.Next() {
		var c database.Competition
		err = rows.Scan(&c.CompetitionId, &c.CompetitionName, &c.Description, &c.StartTime, &c.LeftTime)
		if err != nil {
			return competitions, err
		}
		competitions = append(competitions, c)
	}

	return competitions, nil
}

package models

type Scores struct {
	ScoreId       int64 `json:"score_id"`
	CompetitionId int64 `json:"competition_id"`
	GroupId       int64 `json:"group_id"`
	TotalScore    int   `json:"total_score"`
}

package models

type Scores struct {
	ScoreId       int64 `mapstructure:"score_id"`
	CompetitionId int64 `mapstructure:"competition_id"`
	GroupId       int64 `mapstructure:"group_id"`
	TotalScore    int   `mapstructure:"total_score"`
}

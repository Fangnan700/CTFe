package models

type Solved struct {
	SolvedId      int64 `mapstructure:"solved_id"`
	SolvedTime    int64 `mapstructure:"solved_time"`
	SolvedScore   int   `mapstructure:"solved_score"`
	CompetitionId int64 `mapstructure:"competition_id"`
	ChallengeId   int64 `mapstructure:"challenge_id"`
	GroupId       int64 `mapstructure:"group_id"`
	UserId        int64 `mapstructure:"user_id"`
}

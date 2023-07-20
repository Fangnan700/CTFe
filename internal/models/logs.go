package models

type Logs struct {
	LogId         int64  `mapstructure:"log_id"`
	LogTime       int64  `mapstructure:"log_time"`
	LogType       string `mapstructure:"log_type"`
	LogContent    string `mapstructure:"log_content"`
	CompetitionId int64  `mapstructure:"competition_id"`
	ChallengeId   int64  `mapstructure:"challenge_id"`
	GroupId       int64  `mapstructure:"group_id"`
	UserId        int64  `mapstructure:"user_id"`
}

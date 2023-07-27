package models

type Logs struct {
	LogId         int64  `json:"log_id"`
	LogTime       int64  `json:"log_time"`
	LogType       string `json:"log_type"`
	LogContent    string `json:"log_content"`
	CompetitionId int64  `json:"competition_id"`
	ChallengeId   int64  `json:"challenge_id"`
	GroupId       int64  `json:"group_id"`
	UserId        int64  `json:"user_id"`
}

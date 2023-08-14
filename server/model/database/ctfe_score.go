package database

type Score struct {
	ScoreId       int64 `json:"score_id"`
	CompetitionId int64 `json:"competition_id"`
	ChallengeId   int64 `json:"challenge_id"`
	GroupId       int64 `json:"group_id"`
	UserId        int64 `json:"user_id"`
	Score         int64 `json:"total_score"`
}

package models

type Solved struct {
	SolvedId      int64 `json:"solved_id"`
	SolvedTime    int64 `json:"solved_time"`
	SolvedScore   int   `json:"solved_score"`
	CompetitionId int64 `json:"competition_id"`
	ChallengeId   int64 `json:"challenge_id"`
	GroupId       int64 `json:"group_id"`
	UserId        int64 `json:"user_id"`
}

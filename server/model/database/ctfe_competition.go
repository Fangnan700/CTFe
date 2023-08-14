package database

type Competition struct {
	CompetitionId   int64  `json:"competition_id"`
	CompetitionName string `json:"competition_name"`
	Description     string `json:"description"`
	StartTime       int64  `json:"start_time"`
	LeftTime        int64  `json:"left_time"`
}

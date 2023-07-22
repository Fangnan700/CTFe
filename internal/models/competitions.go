package models

type Competitions struct {
	CompetitionId   int64  `json:"competition_id"`
	CompetitionName string `json:"competition_name"`
	Description     string `json:"description"`
	StartTime       int64  `json:"start_time"`
	LifeTime        int64  `json:"life_time"`
}

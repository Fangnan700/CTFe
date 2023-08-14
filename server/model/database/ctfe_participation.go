package database

type Participation struct {
	ParticipationId int64       `json:"participation_id"`
	GroupId         int64       `json:"group_id"`
	UserId          interface{} `json:"user_id"`
	CompetitionId   interface{} `json:"competition_id"`
	IsAdmin         bool        `json:"is_admin"`
}

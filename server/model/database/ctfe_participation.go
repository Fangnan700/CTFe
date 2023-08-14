package database

type Participation struct {
	ParticipationId int64 `json:"participation_id"`
	GroupId         int64 `json:"group_id"`
	UserId          int64 `json:"user_id"`
	IsLeader        bool  `json:"is_leader"`
}

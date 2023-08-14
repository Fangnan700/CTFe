package database

type Group struct {
	GroupId       int64  `json:"group_id"`
	GroupName     string `json:"group_name"`
	GroupIntro    string `json:"group_intro"`
	CompetitionId int64  `json:"competition_id"`
}

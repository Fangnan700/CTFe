package models

type Groups struct {
	GroupId       int64  `json:"group_id"`
	GroupName     string `json:"group_name"`
	GroupIntro    string `json:"group_into"`
	CompetitionId int64  `json:"competition_id"`
}

package models

type Groups struct {
	GroupId       int64  `mapstructure:"group_id"`
	GroupName     string `mapstructure:"group_name"`
	GroupIntro    string `mapstructure:"group_into"`
	CompetitionId int64  `mapstructure:"competition_id"`
}

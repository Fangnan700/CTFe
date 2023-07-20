package models

type Participation struct {
	GroupId int64 `mapstructure:"group_id"`
	UserId  int64 `mapstructure:"user_id"`
}

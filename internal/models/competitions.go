package models

type Competitions struct {
	CompetitionId   int64  `mapstructure:"competition_id"`
	CompetitionName string `mapstructure:"competition_name"`
	Description     string `mapstructure:"description"`
	StartTime       int64  `mapstructure:"start_time"`
	LifeTime        int64  `mapstructure:"life_time"`
}

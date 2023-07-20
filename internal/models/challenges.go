package models

type Challenges struct {
	ChallengeId   int64  `mapstructure:"challenge_id"`
	ChallengeName string `mapstructure:"challenge_name"`
	ChallengeType string `mapstructure:"challenge_type"`
	Description   string `mapstructure:"description"`
	InitScore     int    `mapstructure:"init_score"`
	ImageName     string `mapstructure:"image_name"`
}

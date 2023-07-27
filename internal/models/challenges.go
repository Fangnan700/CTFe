package models

type Challenges struct {
	ChallengeId   int64  `json:"challenge_id"`
	ChallengeName string `json:"challenge_name"`
	ChallengeType string `json:"challenge_type"`
	Description   string `json:"description"`
	InitScore     int    `json:"init_score"`
	ImageName     string `json:"image_name"`
}

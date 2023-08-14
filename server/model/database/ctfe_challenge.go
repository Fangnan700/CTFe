package database

type Challenge struct {
	ChallengeId      int64  `json:"challenge_id"`
	ChallengeName    string `json:"challenge_name"`
	ChallengeType    string `json:"challenge_type"`
	Description      string `json:"description"`
	ImageName        string `json:"image_name"`
	InitScore        int64  `json:"init_score"`
	DynamicContainer bool   `json:"dynamic_container"`
	DynamicFlag      bool   `json:"dynamic_flag"`
	CompetitionId    int64  `json:"competition_id"`
}

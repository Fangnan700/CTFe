package database

type Container struct {
	ContainerId   int64  `json:"container_id"`
	VesselTag     string `json:"vessel_tag"`
	ContainerHost string `json:"container_host"`
	ContainerPort int    `json:"container_port"`
	StartTime     int64  `json:"start_time"`
	LifeTime      int64  `json:"life_time"`
	CompetitionId int64  `json:"competition_id"`
	ChallengeId   int64  `json:"challenge_id"`
	GroupId       int64  `json:"group_id"`
	UserId        int64  `json:"user_id"`
}

package models

type Containers struct {
	ContainerId   int64  `json:"container_id"`
	VesselId      string `json:"vessel_id"`
	ContainerIp   string `json:"container_ip"`
	ContainerPort int    `json:"container_port"`
	StartTime     int64  `json:"start_time"`
	LifeTime      int64  `json:"life_time"`
	ChallengeId   int64  `json:"challenge_id"`
	GroupId       int64  `json:"group_id"`
}

package models

type Containers struct {
	ContainerId   int64  `mapstructure:"container_id"`
	VesselId      string `mapstructure:"vessel_id"`
	ContainerIp   string `mapstructure:"container_ip"`
	ContainerPort int    `mapstructure:"container_port"`
	StartTime     int64  `mapstructure:"start_time"`
	LifeTime      int64  `mapstructure:"life_time"`
	ChallengeId   int64  `mapstructure:"challenge_id"`
	GroupId       int64  `mapstructure:"group_id"`
}

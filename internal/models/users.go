package models

type Users struct {
	UserId     int64  `mapstructure:"user_id"`
	UserName   string `mapstructure:"user_name"`
	UserSex    string `mapstructure:"user_sex"`
	Email      string `mapstructure:"email"`
	Phone      string `mapstructure:"phone"`
	School     string `mapstructure:"school"`
	StudentNum string `mapstructure:"student_num"`
	CreateTime int64  `mapstructure:"create_time"`
}

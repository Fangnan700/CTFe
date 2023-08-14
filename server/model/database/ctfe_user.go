package database

type User struct {
	UUID       string `json:"uuid"`
	UserId     int64  `json:"user_id"`
	UserName   string `json:"user_name"`
	UserPwd    string `json:"user_pwd"`
	UserSex    string `json:"user_sex"`
	UserEmail  string `json:"user_email"`
	UserPhone  string `json:"user_phone"`
	UserSchool string `json:"user_school"`
	CreateTime int64  `json:"create_time"`
}

package models

type Users struct {
	UserId     int64  `json:"user_id"`
	UserName   string `json:"user_name"`
	UserPwd    string `json:"user_pwd"`
	UserSex    string `json:"user_sex"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	School     string `json:"school"`
	StudentNum string `json:"student_num"`
	CreateTime int64  `json:"create_time"`
}

package database

type Admin struct {
	AdminId int64 `json:"admin_id"`
	UserId  int64 `json:"user_id"`
}

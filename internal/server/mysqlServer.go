package server

import (
	"CTFe/internal/global/mysql"
	"CTFe/internal/models"
)

// SelectAllUsers 查询所有用户
func SelectAllUsers() ([]models.Users, error) {
	return mysql.SelectAllUsers()
}

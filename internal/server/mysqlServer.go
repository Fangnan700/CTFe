package server

import (
	"CTFe/internal/global/mysql"
	"CTFe/internal/models"
)

// SelectLastUserId 查询最后一个用户ID
func SelectLastUserId() (int64, error) {
	return mysql.SelectLastUserId()
}

// SelectAllUsers 查询所有用户
func SelectAllUsers() ([]models.Users, error) {
	return mysql.SelectAllUsers()
}

// InsertUser 添加用户
func InsertUser(registerInfo models.Users) error {
	return mysql.InsertUser(registerInfo)
}

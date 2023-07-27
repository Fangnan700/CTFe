package server

import (
	"CTFe/internal/global/mysql"
	"CTFe/internal/models"
)

// InsertUser 添加用户
func InsertUser(registerInfo models.Users) error {
	return mysql.InsertUser(registerInfo)
}

// SelectLastUserId 查询最后一个用户ID
func SelectLastUserId() (int64, error) {
	return mysql.SelectLastUserId()
}

// SelectAllUsers 查询所有用户
func SelectAllUsers() ([]models.Users, error) {
	return mysql.SelectAllUsers()
}

// SelectUserById 根据用户ID查询用户
func SelectUserById(userId int64) (models.Users, error) {
	return mysql.SelectUserById(userId)
}

// SelectUserByEmailOrPhone 根据用户邮箱/手机查询用户
func SelectUserByEmailOrPhone(email string, phone string) (models.Users, error) {
	return mysql.SelectUserByEmailOrPhone(email, phone)
}

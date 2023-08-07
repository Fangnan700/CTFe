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

// SelectUsers 查询用户
func SelectUsers(keyword interface{}) ([]models.Users, error) {
	return mysql.SelectUsers(keyword)
}

// SelectAllUsers 查询所有用户
func SelectAllUsers() ([]models.Users, error) {
	return mysql.SelectAllUsers()
}

// SelectUserById 根据用户id查询
func SelectUserById(userId interface{}) (models.Users, error) {
	return mysql.SelectUserById(userId)
}

// SelectUserByEmailOrPhone 根据用户邮箱/手机查询用户
func SelectUserByEmailOrPhone(email interface{}, phone interface{}) (models.Users, error) {
	return mysql.SelectUserByEmailOrPhone(email, phone)
}

// UpdateUser 更新用户信息
func UpdateUser(user models.Users) error {
	return mysql.UpdateUser(user)
}

// DeleteUser 删除用户
func DeleteUser(userId interface{}) error {
	return mysql.DeleteUser(userId)
}

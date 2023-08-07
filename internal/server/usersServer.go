package server

import (
	"CTFe/internal/models"
	"CTFe/internal/utils"
	"fmt"
)

// UserRegisterServer 用户注册
func UserRegisterServer(registerInfo models.Users) *models.CTFeError {
	var (
		err     error
		lastId  int64
		preUser models.Users
	)

	// 检查参数是否非空
	if registerInfo.UserName == "" {
		return models.NewCTFeError(0, "用户名不能为空", nil)
	}
	if registerInfo.UserPwd == "" {
		return models.NewCTFeError(0, "密码不能为空", nil)
	}
	if registerInfo.Email == "" {
		return models.NewCTFeError(0, "邮箱不能为空", nil)
	}
	if registerInfo.Phone == "" {
		return models.NewCTFeError(0, "手机号不能为空", nil)
	}

	// 检查邮箱、手机是否已经存在
	preUser, _ = SelectUserByEmailOrPhone(registerInfo.Email, registerInfo.Phone)
	if preUser.UserId != 0 {
		return models.NewCTFeError(-2, "账号已注册", nil)
	}

	// 检查两次密码是否一致
	if registerInfo.UserPwd != registerInfo.UserPwd2 {
		return models.NewCTFeError(-2, "两次输入的密码不一致", nil)
	}

	lastId, _ = SelectLastUserId()
	registerInfo.UserId = lastId + 1
	registerInfo.CreateTime = utils.GetTimeStamp()

	err = InsertUser(registerInfo)
	if err != nil {
		return models.NewCTFeError(-1, "其它异常", err)
	}

	return nil
}

// UserLoginServer 用户登录
func UserLoginServer(loginInfo models.Users) *models.CTFeError {
	var (
		user models.Users
		err  error
	)

	if loginInfo.Email == "" && loginInfo.Phone == "" {
		return models.NewCTFeError(0, "参数不能为空", nil)
	}

	user, err = SelectUserByEmailOrPhone(loginInfo.Email, loginInfo.Phone)
	if err != nil {
		return models.NewCTFeError(-1, "其它异常", err)
	}
	if user.UserId == 0 {
		return models.NewCTFeError(-3, "用户不存在", nil)
	}
	if user.UserPwd != loginInfo.UserPwd {
		return models.NewCTFeError(-3, "账号或密码错误", nil)
	}

	return nil
}

// QueryUsers 查询用户
func QueryUsers(keyword interface{}) ([]models.Users, *models.CTFeError) {
	var (
		users []models.Users
		err   error
	)

	users, err = SelectUsers(keyword)
	if err != nil {
		return users, models.NewCTFeError(-1, "其它异常", err)
	}

	for index, _ := range users {
		users[index].UserPwd = ""
		users[index].UserPwd2 = ""
	}

	return users, nil
}

// QueryAllUsersServer 查询所有用户
func QueryAllUsersServer() ([]models.Users, *models.CTFeError) {
	var (
		users []models.Users
		err   error
	)

	users, err = SelectAllUsers()
	if err != nil {
		return users, models.NewCTFeError(-1, "其它异常", err)
	}

	for index, _ := range users {
		users[index].UserPwd = ""
		users[index].UserPwd2 = ""
	}

	return users, nil
}

// DeleteUserServer 删除用户
func DeleteUserServer(userId interface{}, userPwd interface{}) *models.CTFeError {

	user, err := SelectUserById(userId)
	if err != nil {
		fmt.Println(err)
		return models.NewCTFeError(-1, "其它异常", err)
	}

	if userPwd != user.UserPwd {
		return models.NewCTFeError(-5, "密码错误", nil)
	}

	err = DeleteUser(userId)
	if err != nil {
		return models.NewCTFeError(-1, "其它异常", err)
	}
	return nil
}

// UpdateUserServer 更新用户
func UpdateUserServer(user models.Users) *models.CTFeError {
	var err error

	// 检查参数是否非空
	if user.UserName == "" {
		return models.NewCTFeError(0, "用户名不能为空", nil)
	}
	if user.UserPwd == "" {
		return models.NewCTFeError(0, "密码不能为空", nil)
	}
	if user.Email == "" {
		return models.NewCTFeError(0, "邮箱不能为空", nil)
	}
	if user.Phone == "" {
		return models.NewCTFeError(0, "手机号不能为空", nil)
	}

	// 检查两次密码是否一致
	if user.UserPwd != user.UserPwd2 {
		return models.NewCTFeError(-2, "两次输入的密码不一致", nil)
	}

	err = UpdateUser(user)
	if err != nil {
		return models.NewCTFeError(-1, "其它异常", err)
	}

	return nil

}

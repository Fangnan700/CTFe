package server

import (
	"CTFe/internal/models"
	"CTFe/internal/utils"
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
	registerInfo.UserPwd = utils.MD5Encrypt(registerInfo.UserPwd)
	registerInfo.CreateTime = utils.GetTimeStamp()

	err = InsertUser(registerInfo)
	if err != nil {
		return models.NewCTFeError(-1, "其它异常", err)
	}

	return nil
}

// UserLoginServer 用户登录
func UserLoginServer(loginInfo models.Users) *models.CTFeError {
	// 检查参数是否为空
	if loginInfo.Email == "" && loginInfo.Phone == "" {
		return models.NewCTFeError(0, "参数不能为空", nil)
	}

	// 查询用户信息
	var user models.Users
	var err error

	user, err = SelectUserByEmailOrPhone(loginInfo.Email, loginInfo.Phone)
	if err != nil {
		return models.NewCTFeError(-1, "其它异常", err)
	}
	if user.UserId == 0 {
		return models.NewCTFeError(-3, "用户不存在", nil)
	}
	if user.UserPwd != utils.MD5Encrypt(loginInfo.UserPwd) {
		return models.NewCTFeError(-3, "账号或密码错误", nil)
	}

	return nil
}

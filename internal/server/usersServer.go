package server

import (
	"CTFe/internal/models"
	"CTFe/internal/utils"
)

// UserRegisterServer 用户注册
func UserRegisterServer(registerInfo models.Users) error {
	var err error
	var lastId int64

	lastId, err = SelectLastUserId()
	if err != nil {
		return err
	}

	registerInfo.UserId = lastId + 1
	registerInfo.CreateTime = utils.GetTimeStamp()

	err = InsertUser(registerInfo)
	if err != nil {
		return err
	}

	return nil
}

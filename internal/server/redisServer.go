package server

import "CTFe/internal/global/redis"

// GetCTFeTokenStatus 获取用户登录状态
func GetCTFeTokenStatus(ctfeToken string) int {
	return redis.GetCTFeTokenStatus(ctfeToken)
}

// SetCTFeTokenStatus 设置用户登录状态
func SetCTFeTokenStatus(ctfeToken string, status int) {
	redis.SetCTFeTokenStatus(ctfeToken, status)
}

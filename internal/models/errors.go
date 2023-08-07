package models

import "fmt"

type CTFeError struct {
	Code    int // 0：输入为空  -1：其它异常  -2：注册异常  -3：登录异常  -4：查询异常  -5：权限异常
	Message string
	Content interface{}
}

func (e *CTFeError) Error() string {
	return fmt.Sprintf("Code:%d\nMessage:%s\nContent:%++v\n", e.Code, e.Message, e.Content)
}

func NewCTFeError(code int, message string, content interface{}) *CTFeError {
	return &CTFeError{
		Code:    code,
		Message: message,
		Content: content,
	}
}

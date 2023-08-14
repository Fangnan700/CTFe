package log

import "testing"

func TestLogger(t *testing.T) {
	InfoLogger.Println("信息日志测试")
	ErrorLogger.Println("错误日志测试")
}

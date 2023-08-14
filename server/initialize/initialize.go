package initialize

import (
	"CTFe/server/util/log"
	"fmt"
)

func Init() {
	PrintCopyright()
	log.InitLogger()
}

func PrintCopyright() {
	fmt.Printf(`

	欢迎使用 CTFe
	当前版本：v1.0.0
	项目地址：https://github.com/Fangnan700/CTFe

`)
	return
}

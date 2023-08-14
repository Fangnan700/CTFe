package main

import (
	"CTFe/server/initialize"
	"CTFe/server/router"
)

func main() {
	initialize.Init()
	router.Start()
}

package main

import (
	"CTFe/internal/routers"
	"fmt"
)

func main() {
	var err error

	err = routers.Register()
	if err != nil {
		fmt.Println(err)
	}
}

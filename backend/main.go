package main

import (
	"github.com/ritscc/Snack/driver"
)

func main() {
	// 1. 8080番portのLisnterを作成
	s := driver.InitServer(8080)

	s.StartServer()

}

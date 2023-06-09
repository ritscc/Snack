package main

import (
  "net"
  "fmt"

	"github.com/ritscc/Snack/driver"
)

func main() {
	// 1. 8080番portのLisnterを作成
	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	driver.StartServer(listener)

}

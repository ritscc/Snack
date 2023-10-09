package main

import (
	"log"

	"github.com/ritscc/Snack/internal/driver"
)

func main() {
	config := driver.Config{}
	if err := driver.LoadConfig(&config); err != nil {
		log.Fatalln(err)
	}

	s := driver.NewServer(config.ServerConfig)

	if err := s.StartServer(); err != nil {
		log.Fatalln(err)
	}

}

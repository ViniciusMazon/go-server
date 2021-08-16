package main

import (
	"github.com/viniciusmazon/go-server/database"
	"github.com/viniciusmazon/go-server/server"
)

func main() {
	database.StartDB()
	s := server.NewServer()

	s.Run()
}

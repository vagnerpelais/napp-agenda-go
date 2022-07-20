package main

import (
	"github.com/vagnerpelais/napp-agenda/config"
	"github.com/vagnerpelais/napp-agenda/database"
	"github.com/vagnerpelais/napp-agenda/server"
)

func main() {
	config.Load()
	database.StartDB(config.ConnectionString)
	s := server.NewServer()
	s.Run()
}

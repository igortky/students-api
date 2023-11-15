package main

import (
	"go_gin/database"
	"go_gin/routes"
)

func main() {
	database.DBConnect()
	routes.HandleRequests()
}

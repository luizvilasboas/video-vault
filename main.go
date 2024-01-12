package main

import (
	"gitlab.com/olooeez/video-vault/database"
	"gitlab.com/olooeez/video-vault/routes"
)

func main() {
	database.Connect()
	routes.HandleRequests()
}

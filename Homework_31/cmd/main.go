package main

import (
	"task31/internal/app"
	mongodb "task31/scripts/client/mongoDB"
)

func main() {
	mongodb.ConnectAndPing()
	app.StartApp() //	"task31/internal/app"

}

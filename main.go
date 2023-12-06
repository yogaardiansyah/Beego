// main.go
package main

import (
	"Bego/config"
	"Bego/controllers/control"
	"Bego/env"
	"Bego/server"
)

func main() {
	// Koneksi Database
	config.ConnectDB()

	// Routing
	// 1. Home Routing
	control.Router("/testing", "home/index.html", "GET")

	// Server Run
	server.RunServer(env.ServerAddr)
}

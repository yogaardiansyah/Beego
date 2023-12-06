package main

import (
	"Beego/config"
	"Beego/controllers/control"
	"log"
	"net/http"
)

func main() {
	// Koneksi Database
	config.ConnectDB()

	// Routing
	// 1. Home Routing
	control.Router("/testing", "home/index.html", "GET")

	// Server Run
	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)

}

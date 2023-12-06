package main

import (
	"Beego/config"
	"Beego/controllers/controlhome"
	"log"
	"net/http"
)

func main() {
	// Koneksi Database
	config.ConnectDB()

	// Routing
	// 1. Home Routing
	http.HandleFunc("/", controlhome.Welcome)

	// Server Run
	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)

}

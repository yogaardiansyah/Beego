package main

import (
	"Beego/config"
	"Beego/controllers/control"
	"log"
	"net/http"
	"path/filepath"
	"runtime"
)

func main() {
	// Koneksi Database
	config.ConnectDB()

	// Routing
	// 1. Home Routing
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)

	http.HandleFunc("/test", control.Router(basePath, "home/index.html"))

	// Server Run
	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)

}

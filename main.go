package main

import (
	"Bego/config"
	"Bego/controllers/control"
	"log"
	"net"
	"net/http"
	"strings"
)

func main() {
	// Koneksi Database
	config.ConnectDB()

	// Routing
	// 1. Home Routing
	control.Router("/testing", "home/index.html", "GET")

	// Server Run
	server := &http.Server{Addr: ":8080"}
	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	// Get the local address after the server starts
	addr := server.Addr
	if addr == "" {
		addr = ":http"
	}

	// Get the IP address of the server
	ip, err := getLocalIP()
	if err == nil {
		addr = ip + addr[strings.LastIndex(addr, ":"):]
	}

	log.Printf("Server running on http://%s", addr)

	// Allow the server to keep running until manually stopped (Ctrl+C)
	select {}
}

func getLocalIP() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80") // Google's public DNS IP
	if err != nil {
		return "", err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String(), nil
}

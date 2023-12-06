// server.go
package server

import (
	"log"
	"net"
	"net/http"
	"strings"
)

func RunServer(addr string) {
	server := &http.Server{Addr: addr}
	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	// Get the local address after the server starts
	localAddr := server.Addr
	if localAddr == "" {
		localAddr = ":http"
	}

	// Get the IP address of the server
	ip, err := getLocalIP()
	if err == nil {
		localAddr = ip + localAddr[strings.LastIndex(localAddr, ":"):]
	}

	log.Printf("Server running on http://%s", localAddr)

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

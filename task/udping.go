package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// Replace with your server IP and port
	server := "2606:4700:100::/48:2048"

	conn, err := net.DialTimeout("tcp", server, 5*time.Second)
	if err != nil {
		fmt.Printf("Failed to connect to server: %v\n", err)
		return
	}
	defer conn.Close()

	start := time.Now()
	_, err = conn.Write([]byte("ping"))
	if err != nil {
		fmt.Printf("Failed to send data to server: %v\n", err)
		return
	}

	buf := make([]byte, 1024)
	_, err = conn.Read(buf)
	if err != nil {
		fmt.Printf("Failed to read data from server: %v\n", err)
		return
	}

	elapsed := time.Since(start)
	fmt.Printf("Round trip time: %v\n", elapsed)
}

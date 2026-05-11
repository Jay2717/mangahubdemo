package udp

import (
	"fmt"
	"net"
)

func StartUDPServer() {
	addr, err := net.ResolveUDPAddr("udp", ":6060")
	if err != nil {
		panic(err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	fmt.Println("UDP Server running on :6060")

	buffer := make([]byte, 1024)

	for {
		n, remoteAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			continue
		}

		message := string(buffer[:n])

		fmt.Println("=== UDP Notification ===")
		fmt.Println("From:", remoteAddr)
		fmt.Println("Message:", message)
	}
}
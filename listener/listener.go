package listener

import (
	"fmt"
	"hw-api/publisher"
	"net"
)

func Listener() {
	// Listen for incoming connections on port 8080
	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Accept incoming connections and handle them
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Handle the connection in a new goroutine
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	// Close the connection when we're done
	defer conn.Close()

	// Read incoming data
	buf := make([]byte, 12)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	packet := publisher.Decode(buf)

	// Print the incoming data
	fmt.Printf(
		"received new data: sid: %d, locID %d ts: %d, temp:%d\n",
		packet.Sensid, packet.Locid, packet.Tstamp, packet.Temp,
	)
}

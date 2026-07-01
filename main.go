package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
)

// type person struct {
// 	Name string
// 	Age int
// }

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Failed to bind port 8080:", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("🚀 Core network engine online on http://localhost:8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go func(c net.Conn) {
			defer c.Close()

			// var buf []byte

			// for c.Read() {

			// }

			var requestData []byte
			for {
				buf := make([]byte, 16)
				n, err := c.Read(buf)
				requestData = append(requestData, buf[:n]...)

				if err != nil {
					return
				}

				if bytes.Contains(requestData, []byte("\r\n\r\n")) {
					break
				}

			}
			fmt.Printf("📥 Received Request:\n%s\n", string(requestData))

			response := "HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: 17\r\n\r\nHello from Bhuvi!" // \r\nContent-Type: text/plain\r\nContent-Length: 20\r\n\r\nHello from Bhuvnesh!"
			// response := ""
			c.Write([]byte(response))
		}(conn)
	}

}

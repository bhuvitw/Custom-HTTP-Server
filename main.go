package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"strings"
)

type HTTPRequest struct {
	Method  string
	Path    string
	Version string
	Headers map[string]string
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")

	if err != nil {
		fmt.Println("Error: ", err)
	}

	defer listener.Close()
	fmt.Println("localhost:8080 online")

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("Error: ", err)
		}

		var requestedarr []byte
		for {
			buf := make([]byte, 16)
			len, err := conn.Read(buf)
			if err != nil {
				fmt.Println("Error: ", err)
			}
			requestedarr = append(requestedarr, buf[:len]...)

			if bytes.Contains(requestedarr, []byte("\r\n\r\n")) {
				break
			}
		}

		fmt.Printf("Recieved: Request: \n%s\n", string(requestedarr))

		m, err := ParseRequest(requestedarr)

		var response []byte

		if m.Path == "/profile.html" {
			response = []byte("HTTP/1.1 200 OK\r\n\r\nWelcome to Bhuvi's Profile\n")
		} else if m.Path == "/" {
			response = []byte("HTTP/1.1 200 OK\r\n\r\nHello World\n")
		} else if m.Path == "/index.html" {
			content, err := os.ReadFile("index.html")
			if err != nil {
				fmt.Println("Error Reading index.html", err)

				response = []byte("HTTP/1.1 404 Internarl Server Error\r\n\r\nCould not load File")
			} else {
				headerString := fmt.Sprintf(
					"HTTP/1.1 200 OK\r\nContent-Type:text/html\r\nContent-Length: %d\r\n\r\n",
					len(content),
				)

				response = append([]byte(headerString), content...)
			}

		} else if m.Path == "/quote.png" {
			content, err := os.ReadFile("quote.png")
			if err != nil {
				fmt.Println("Error Reading quote.png", err)

				response = []byte("HTTP/1.1 404 Internal Server Error\r\n\r\nCount not load image")
			} else {
				headerString := fmt.Sprintf(
					"HTTP/1.1 200 OK\r\nContent-Type:image/png\r\nContent-Length: %d\r\n\r\n",
					len(content),
				)

				response = append([]byte(headerString), content...)
			}

		} else {
			response = []byte("HTTP/1.1 404 Not Found\r\n\r\nPage Not Found\n")
		}

		conn.Write(response)

		conn.Close()

	}

}

func ParseRequest(rawData []byte) (*HTTPRequest, error) {
	data := strings.Split(string(rawData), "\r\n")

	firstLine := strings.Split(string(data[0]), " ")

	var self HTTPRequest

	self.Method = firstLine[0]
	self.Path = firstLine[1]
	self.Version = firstLine[2]

	header := make(map[string]string)

	for i := 1; i < len(data); i++ {
		if data[i] == "" {
			break
		}
		saver(header, string(data[i]))
	}

	self.Headers = header

	return &self, nil

}

func saver(header map[string]string, content string) {
	ans := strings.SplitN(content, ":", 2)
	header[ans[0]] = ans[1]
}

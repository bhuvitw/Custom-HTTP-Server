package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime"
	"net"
	"os"
	"path/filepath"
	"strings"
)

type HTTPRequest struct {
	Method  string
	Path    string
	Version string
	Headers map[string]string
	Body    []byte
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
		buf := make([]byte, 16)
		for {
			n, err := conn.Read(buf)
			if err != nil {
				if errors.Is(err, io.EOF) {
					fmt.Println("Client Closed Connection")
				} else {
					fmt.Println("Error: ", err)
				}
				break
			}
			requestedarr = append(requestedarr, buf[:n]...)

			if len(requestedarr) > 8192 {
				break
			}
			if bytes.Contains(requestedarr, []byte("\r\n\r\n")) {

			}
		}

		fmt.Printf("Recieved: Request: \n%s\n", string(requestedarr))

		m, err := ParseRequest(requestedarr)

		if err != nil {
			fmt.Println("Error: ", err)
			conn.Close()
			continue
		}

		response := Router(m)

		conn.Write(response)

		conn.Close()

	}

}

func Router(m *HTTPRequest) []byte {

	if m.Path == "/" {
		body := []byte("Hello World\n")
		header := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n", len(body))
		return append([]byte(header), body...)
	}

	if m.Method == "POST" && m.Path == "/echo" {
		body := m.Body
		header := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContnet-Length: %d\r\n\r\n", len(body))
		return append([]byte(header), body...)
	}
	filePath := strings.TrimPrefix(m.Path, "/")

	content, err := os.ReadFile(filePath)

	if err != nil {
		if os.IsNotExist(err) {
			body := []byte("404 Page Not Found")
			header := fmt.Sprintf("HTTP/1.1 404 Not Found\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n", len(body))
			return append([]byte(header), body...)
		}

		body := []byte("500 Internal Server Error")

		header := fmt.Sprintf("HTTP/1.1 500 Internal Server Error\r\n\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n", len(body))
		return append([]byte(header), body...)
	}

	ext := filepath.Ext(filePath)

	contentType := mime.TypeByExtension(ext)
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	headerString := fmt.Sprintf(
		"HTTP/1.1 200 Ok\r\n"+
			"Content-Type: %s\r\n"+
			"Content-Length: %d\r\n"+
			"Connection: close\r\n\r\n",
		contentType,
		len(content),
	)

	response := append([]byte(headerString), content...)

	return response

}

func ParseRequest(rawData []byte) (*HTTPRequest, error) {
	data := strings.Split(string(rawData), "\r\n\r\n")
	Head := strings.Split(string(data[0]), "\r\n")

	firstLine := strings.Split(string(Head[0]), " ")

	if len(firstLine) != 3 {
		return nil, errors.New("first line length is not correct")
	}

	var self HTTPRequest

	self.Method = firstLine[0]
	self.Path = firstLine[1]
	self.Version = firstLine[2]
	self.Body = []byte(data[1])

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
	header[strings.TrimSpace(strings.ToLower(ans[0]))] = strings.TrimSpace(ans[1])
}

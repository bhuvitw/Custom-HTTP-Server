package main

import (
	"bytes"
	"fmt"
	"net"
	"strings"
)

func main() {
	// spliter("GET /profile.html HTTP/1.1")

	// header := make(map[string]string)
	// saver(header, "Content-Type: text/html")

	listener, err := net.Listen("tcp", "localhost:8080")

	if err != nil {
		fmt.Println("Error: ", err)
	}

	defer listener.Close()
	fmt.Println("localhost:8080 online")

	for {
		header := make(map[string]string)
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
		data := strings.Split(string(requestedarr), "\r\n")
		//  firstLine:= spliter(string(data[0]))
		firstLine := strings.Split(string(data[0]), " ")
		for i := 1; i < len(data); i++ {
			if data[i] == "" {
				break
			}
			saver(header, string(data[i]))
		}
		fmt.Println("firstLine: ", firstLine)
		fmt.Println("map: ", header)

		var response []byte

		if firstLine[1] == "/profile.html" {
			response = []byte("HTTP/1.1 200 OK\r\n\r\nWelcome to Bhuvi's Profile\n")
		} else if firstLine[1] == "/" {
			response = []byte("HTTP/1.1 200 OK\r\n\r\nHello World\n")
		} else {
			response = []byte("HTTP/1.1 404 Not Found\r\n\r\nPage Not Found\n")
		}

		conn.Write(response)

		conn.Close()

	}

}

func saver(header map[string]string, content string) {
	ans := strings.SplitN(content, ":", 2)
	header[ans[0]] = ans[1]

	// fmt.Println(header)
}

func spliter(content string) {
	// fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Println("%q\n ", strings.Split(content, " "))
}

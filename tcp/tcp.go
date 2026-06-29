package main

import (
	"fmt"
	"net"
)

func createServer() {
	// conn, err := net.Dial("tcp", "http://localhost:8080/")
	// if err != nil {
	// 	fmt.Println("Server not created")
	// }
	// fmt.Fprintf(conn, "Get / HTTP/1.0\r\n\r\n")

	// code to create a server using listen
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
		fmt.Println(err, "error creating server")
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error recieving maybe")
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	// conn.Read([]byte bitu)
	// fmt.Println(bitu)
	// // learning how to create slice in go
	// letters := []string{"a", "b", "c", "d"}
	// s = make([]byte, 5, 5)
	s := [1024]byte{}
	l, err := conn.Read(s[:])
	if err != nil {
		fmt.Println(err, "Error")
	}
	fmt.Println(s[:], l)

}

// func connectServer(){
// 	ln, err :=
// }

func main() {
	createServer()

	// // code to connect to the server
	// conn, err := net.Dial("tcp", "localhost:8080")
	// if err != nil {
	// 	// handle error
	// }
	// fmt.Fprintf(conn, "Get / HTTP/1.0\r\n\r\n")
	// status, err := bufio.NewReader(conn).ReadString('\n')
	// //..

	// // code to create a server using listen
	// ln, err := net.Listen("tcp", ":8080")
	// if err != nil {
	// 	// handle error
	// 	fmt.Println("error creating server")
	// }
	// for {
	// 	conn, err := ln.Accept()
	// 	if err != nil {
	// 		fmt.Println("Faield to connect to the server")
	// 		go handleConnection(conn)
	// 	}

	// }

}

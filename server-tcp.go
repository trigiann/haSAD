package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	server, err := net.Listen("tcp", ":7000")
	fmt.Println("server listening on 7000")
	if err != nil {
		log.Fatalln(err)
	}
	defer server.Close()

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("new connection")

		go listenConnection(conn)
	}
}

func listenConnection(conn net.Conn) {
	for {
		buffer := make([]byte, 1400)
                dataSize, err := conn.Read(buffer)
                if err != nil {
                        fmt.Println("connection closed")
                        return
                }
               
                data := buffer[:dataSize]
                fmt.Println("received message: ", string(data))
               
                _, err = conn.Write(data)
                if err != nil {
                        log.Fatalln(err)
                }
                fmt.Println("Message sent: ", string(data))
	}
}

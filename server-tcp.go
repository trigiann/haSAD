package main

import (
	"fmt"
	"log"
	"net"
)

var (
	clients = make(map[string]net.Conn)
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
		go listenConnection(conn)

	}

}

func listenConnection(conn net.Conn) {
	buffer := make([]byte, 1400)
	bufferNick := make([]byte, 1400)
	dataSizeNick, _ := conn.Read(bufferNick)
	nick := bufferNick[:dataSizeNick]
	for key, _ := range clients{
		for string(nick) == key{
			conn.Write([]byte("Username already in use. Choose another nick: "))
			dataSizeNick, _ = conn.Read(bufferNick)
			nick = bufferNick[:dataSizeNick]
		}  
	}
	conn.Write([]byte("Welcome to the chat!"))
	//Register NEW client
	clients[string(nick)] = conn

	for {
		dataSize, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("connection closed")
			return
		}

		data := buffer[:dataSize]

		response := string(nick) + ": " + string(data)
		for _, i := range clients {

			_, err = i.Write([]byte(response))
			if err != nil {
				log.Fatalln(err)
			}

		}

	}

}

package main

import (
	"fmt"
	"log"
	"net"
)

//var (clients = make (map[net.Conn]string))

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
		fmt.Println("new connection: ")
	}
		
}


func listenConnection(conn net.Conn) {
	// buffer := make([]byte, 1400)
 //    dataSize, err := conn.Read(buffer)
 //    if err != nil {
 //            fmt.Println("connection closed")
 //            return
 //    }
 //    nick := buffer[:dataSize]
    

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


/*func listenConnection(conn net.Conn) {
	buffer := make([]byte, 1400)
 	dataSize,_ := conn.Read(buffer)
    nick := buffer[:dataSize]
    
	//Register NEW client
	clients[conn] = string(nick)

	for clients := range clients {
		dataSize, err := conn.Read(buffer)
        if err != nil {
                fmt.Println("connection closed")
                return
        }
       
        data := buffer[:dataSize]
        fmt.Println("received message: ", string(data))
       
        _, err = clients.Write(data)
        if err != nil {
                log.Fatalln(err)
        }
        fmt.Println("Message sent: ", string(data))


	}


}*/
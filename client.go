package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)

	fmt.Print("Choose a nickname: ")
	r := bufio.NewReader(os.Stdin)
	nick, _, _ := r.ReadLine()
	conn, err := net.Dial("tcp", ":7000")
	_, _ = conn.Write(nick)

	if err != nil {
		log.Fatalln(err)
	}

	go keyboard(conn)
	go readServer(conn)
	wg.Wait()

}

func keyboard(conn net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	for data != nil {
		_, err := conn.Write(data)
		if err != nil {
			log.Fatalln(err)
		}
		data, _, _ = reader.ReadLine()
	}
	wg.Done()
}

func readServer(conn net.Conn) {
	buffer := make([]byte, 1400)
	dataSize, err := conn.Read(buffer)
	info := buffer[:dataSize]
	for info != nil {
		if err != nil {
			fmt.Println("connection closed")
			return
		}

		fmt.Println(string(info))
		dataSize, _ = conn.Read(buffer)
		info = buffer[:dataSize]
	}
	wg.Done()
}

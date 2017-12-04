package main
 
import (
        "fmt"
        "log"
        "net"
)
 
func main() {
        conn, err := net.Dial("tcp", ":7000")
        if err != nil {
                log.Fatalln(err)
        }
 
        _, err = conn.Write([]byte("Hello Server!"))
        if err != nil {
                log.Fatalln(err)
        }
        fmt.Println("Hem saludat al server")
 
        _, err = conn.Write([]byte("How are you?"))
        if err != nil {
                log.Fatalln(err)
        }
        fmt.Println("Hem preguntat com esta el server")
 
        for {
                buffer := make([]byte, 1400)
                dataSize, err := conn.Read(buffer)
                if err != nil {
                        fmt.Println("connection closed")
                        return
                }
 
                data := buffer[:dataSize]
                fmt.Println("received message: ", string(data))
        }
}
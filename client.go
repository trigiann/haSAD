package main
 
import (
        "fmt"
        "log"
        "net"
        "bufio"
        "os"
)
 
func main() {
        // fmt.Print("Choose a nickname: ")
        // r := bufio.NewReader(os.Stdin)
        // nick,_,_ := r.ReadLine()
        // conn, err := net.Dial("tcp", ":7000")
        // _,_ = conn.Write(nick)

        conn, err := net.Dial("tcp", ":7000")

        if err != nil {
                log.Fatalln(err)
        }
        for {
                data := read()
                _, err = conn.Write(data)
                if err != nil {
                        log.Fatalln(err)
                }
                buffer := make([]byte, 1400)
                dataSize, err := conn.Read(buffer)
                if err != nil {
                        fmt.Println("connection closed")
                        return
                }
 
                info := buffer[:dataSize]
                fmt.Println("received message: ", string(info))
        }
}

func read () []byte {
        reader := bufio.NewReader(os.Stdin)
        data,_,_ := reader.ReadLine()
        return data
} 
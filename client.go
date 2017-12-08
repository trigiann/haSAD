package main
 
import (
        "fmt"
        "log"
        "net"
        "bufio"
        "os"
)
 
func main() {

         fmt.Print("Choose a nickname: ")
         r := bufio.NewReader(os.Stdin)
         nick,_,_ := r.ReadLine()
         conn, err := net.Dial("tcp", ":7000")
         _,_ = conn.Write(nick)

        //conn, err := net.Dial("tcp", ":7000")

        if err != nil {
                log.Fatalln(err)
        }
        for {
                go keyboard(conn)
                go readServer(conn)
        }
}

func keyboard (conn net.Conn) {
        reader := bufio.NewReader(os.Stdin)
        data,_,_ := reader.ReadLine()
        _, err := conn.Write(data)
        if err != nil {
                log.Fatalln(err)
        }
} 

func readServer (conn net.Conn) {
        buffer := make([]byte, 1400)
        dataSize, err := conn.Read(buffer)
        if err != nil {
                fmt.Println("connection closed")
                return
        }
        info := buffer[:dataSize]
        fmt.Println(string(info))
}
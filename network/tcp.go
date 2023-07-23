package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")

	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g ,connection aborted
			continue
		}
		handleConn(conn)
	}
}

func handleConn(con net.Conn) {

	buff := make([]byte, 1000)
	_, err := con.Read(buff)
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(8 * time.Second)
	resp := "HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\n\r\nHello World to HTTP Server!!\r\n"
	con.Write([]byte(resp))
	fmt.Println(time.Now())
	con.Close()

}

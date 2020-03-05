package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	cancel := make(chan int)
	go func() {
		var mon string
		fmt.Scanln(&mon)
		if mon == "exit" {
			cancel <- 1
		}
	}()

	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		select {
		case <-cancel:
			conn.Close()
			listener.Close()
			break
		default:
			go handleConn(conn)
		}
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n\r"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

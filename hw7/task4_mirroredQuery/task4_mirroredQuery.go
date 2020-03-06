package main

import (
	"fmt"
	"log"
	"net"
)

func mirroredQuery() string {
	responses := make(chan string, 3)
	go func() {
		responses <- request("ya.ru:80") //один из адесов Яндекса
	}()
	go func() {
		responses <- request("www.yandex.ru:80") //один из адесов Яндекса
	}()
	go func() {
		responses <- request("77.88.8.8:80") //один из адесов Яндекса
	}()
	return <-responses // возврат самого быстрого ответа
}

func request(hostname string) string {
	conn, err := net.Dial("tcp", hostname)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	return hostname
}

func main() {
	fmt.Println(mirroredQuery())
}

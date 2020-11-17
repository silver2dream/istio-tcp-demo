package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	host := os.Args[1]
	port := os.Args[2]
	addr := fmt.Sprintf("%v:%v", host, port)
	res, err := startTCP(addr)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

	fmt.Println(res)

	wait := make(chan bool)
	<-wait
}

func startTCP(addr string) (string, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	go func() {
		for {
			conn.Write([]byte(time.Now().String()))
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		bs := make([]byte, 1024)
		len, err := conn.Read(bs)
		if err != nil {
			fmt.Println(err)

		} else {
			fmt.Println(string(bs[:len]), nil)
		}
	}()

	return "Heartbeat Start", nil
}

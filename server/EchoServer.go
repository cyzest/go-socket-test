package main

import "net"
import "fmt"

func echoServer(c net.Conn) {
	for {
		buf := make([]byte, 512)
		nr, err := c.Read(buf)
		if err != nil {
			return
		}
		data := buf[0:nr]
		fmt.Println("Received:", string(data))
		_, err = c.Write(data)
		if err != nil {
			//fmt.Println("Write: " + err.String())
		}
	}
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		//fmt.Println("listen error", err.String())
		return
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			//fmt.Println("accept error", err.String())
			continue
		}
		go echoServer(conn)
	}
}

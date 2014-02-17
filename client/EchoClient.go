package main

import (
	"fmt"
	"io"
	"net"
	//"time"
)

func reader(r io.Reader) {
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf[:])
		if err != nil {
			return
		}
		fmt.Println("Response Message:", string(buf[0:n]))
	}
}

func main() {
	c, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		//fmt.Println(err.String())
	}
	defer c.Close()
	go reader(c)
	for {
		var message string
		//fmt.Print("Request Message:")
		_, err1 := fmt.Scanln(&message)
		if err1 != nil {
			//fmt.Println(err.String())
			continue
		}
		_, err2 := c.Write([]byte(message))
		if err2 != nil {
			//fmt.Println(err.String())
			break
		}
		//time.Sleep(1e9)
	}
}

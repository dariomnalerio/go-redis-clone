package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		panic(err)
	}

	conn, err := l.Accept()

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	for {
		buff := make([]byte, 1024)

		_, err = conn.Read(buff)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("error reading from client: ", err.Error())
			os.Exit(1)
		}

		conn.Write([]byte("+OK\r\n"))
	}
}

package main

import (
	"io"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go func(c net.Conn) {
			readBuf := make([]byte, 100)
			io.WriteString(c, "220 JiffyMail Hello [192.168.1.234]\r\n")
			io.ReadFull(c, readBuf)
			io.WriteString(c, string(readBuf))
			c.Close()
		}(conn)
	}
}

package main

import (
	"fmt"
	"net"
	"time"
	"flag"
	"net/http"
)

func StreamReceive(c chan item) {

	ServerAddr, err := net.ResolveUDPAddr("udp", ":7778")
	_check (err)
	fmt.Println("listening on :7778")

	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	_check (err)
	defer ServerConn.Close()

	buf := make([]byte, 1024)

	for {
		n, addr, err := ServerConn.ReadFromUDP(buf)
		c <- item{string(buf[0:n]), *addr, time.Now()}
		_check (err)
	}
}

func netStarter(port string)  error{
		var addr = flag.String("addr", "localhost:8080", "http service address")
		err := http.ListenAndServe(*addr, nil)
		return err
}

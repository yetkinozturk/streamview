package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"time"
)

func StreamReceive(c chan item, udpPort string) {

	ServerAddr, err := net.ResolveUDPAddr("udp", ":"+udpPort)
	_check(err)
	fmt.Println("listening on :" + udpPort)

	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	_check(err)
	defer ServerConn.Close()

	buf := make([]byte, 1024)

	for {
		n, addr, err := ServerConn.ReadFromUDP(buf)
		c <- item{string(buf[0:n]), *addr, time.Now()}
		_check(err)
	}
}

func netStarter(port string) error {
	var addr = flag.String("addr", "localhost:"+port, "http service address")
	err := http.ListenAndServe(*addr, nil)
	return err
}

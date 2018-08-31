package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"time"
)

func (sv *StreamView) udpStarter() {

	ServerAddr, err := net.ResolveUDPAddr("udp", ":"+sv.udpPort)
	_check(err)

	fmt.Println("listening UDP on :" + sv.udpPort)

	ServerConn, err := net.ListenUDP("udp", ServerAddr)

	defer ServerConn.Close()

	buf := make([]byte, 1024)

	for {
		n, addr, err := ServerConn.ReadFromUDP(buf)
		_check(err)
		sv.netChan <- netMessages{string(buf[0:n]), *addr, time.Now()}
	}
}

func (sv StreamView) httpStarter() error {
	sv.httpRouteConfig()
	var addr = flag.String("addr", "localhost:"+sv.httpPort, "http service address")
	fmt.Println("listening HTTP on :" + sv.httpPort)
	err := http.ListenAndServe(*addr, nil)

	return err
}

func (sv *StreamView) start() error {

	//go func() {
	go sv.udpStarter()
	//}()	
	err := sv.httpStarter()
	

	return err
}

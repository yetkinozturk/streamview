package main

import (
	"log"
	"net"
	"time"
)

type netMessages struct {
	message   string
	addr      net.UDPAddr
	timestamp time.Time
}

type StreamView struct {
	netChan  chan netMessages
	udpPort  string
	httpPort string
}

type SamplePKG struct {
	Power     bool
	Execution bool
	Rand      float64
	HLimit    int32
	values    []float32
}

// _check  Terminate recv on error
func _check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func NewStreamView(udpPort string, httpPort string) *StreamView {
	sv := new(StreamView)
	sv.netChan = make(chan netMessages)
	sv.udpPort = udpPort
	sv.httpPort = httpPort
	return sv
}

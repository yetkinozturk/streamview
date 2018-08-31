package main

import (
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


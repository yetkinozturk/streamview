package main

import "net"
import "time"
import "fmt"
import "os"

type item struct {
	message   string
	addr      net.UDPAddr
	timestamp time.Time
}

type SamplePKG struct {
	Power bool
	Execution bool
	Rand float64
	HLimit int32
	values []float32
}

// _check  Terminate recv on error
func _check(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}
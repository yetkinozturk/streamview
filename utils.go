package streamview

import (
	"log"
)

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

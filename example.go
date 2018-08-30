package main

import (
	"flag"
	"fmt"
	"log"
	"os"

    "github.com/yetkinozturk/streamview"
)

func main() {

	flag.Parse()
	log.SetFlags(0)

	if len(os.Args) != 3 {
		fmt.Println("Usage: ./streamview udpport httpport")
		return
	}

	streamView := NewStreamView(os.Args[1], os.Args[2])
	log.Fatal(streamView.start())

}

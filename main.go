package main

import (
	"fmt"
	"flag"
	"log"
	"os"
)

func main() {
	if len(os.Args) == 0 {
		fmt.Println("Usage: ./streamview udpport httpport")
		return
	}
	//udp_port := os.Args[1]
	http_port := os.Args[2]
	flag.Parse()
	log.SetFlags(0)
	httpRouteConfig()
	log.Fatal(netStarter(http_port))
}


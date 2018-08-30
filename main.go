package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ./streamview udpport httpport")
		return
	}
	udpPort := os.Args[1]
	httpPort := os.Args[2]
	flag.Parse()
	log.SetFlags(0)
	httpRouteConfig(udpPort)
	log.Fatal(netStarter(httpPort))
}

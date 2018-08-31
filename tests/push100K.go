package main

import (
    "fmt"
    "net"
    "time"
)

func main() {
    const limit = 100000
    const addr = "127.0.0.1:7778"

    fmt.Printf("Pushing %v messages to %v\n",limit, addr)
    conn, err := net.Dial("udp", addr)
    if err != nil {
        fmt.Printf("Some error %v", err)
        return
    }
    time1 := time.Now()

    for i:=0; i<limit; i+=1 {
        fmt.Fprintf(conn, "msg: %v",i)
    }
    elapsed := time.Since(time1)
    fmt.Printf("Time elapsed %s\n", elapsed)

    conn.Close()
}
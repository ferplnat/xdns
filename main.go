package main

import (
	"fmt"
	"log"
	"net"
	"time"
	"xdns/config"
)

func main() {
    confPath := ".\\conf.yaml"
    conf, err := config.Load(confPath)
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Listening on: %d", conf.Port)

    port := fmt.Sprintf(":%d", conf.Port)
    ln, err := net.ListenPacket("udp", port)
    if err != nil {
        log.Fatal(err)        
    }
    defer ln.Close()

    for {
        requestBuf := make([]byte, 1024)
        bytesRead, addr, err := ln.ReadFrom(requestBuf)
        if bytesRead < 0 {
        }
        if err != nil {
            fmt.Printf("Error reading: %v", err)            
        }

        log.Printf("Received request: %d bytes.", bytesRead)
        log.Printf(string(requestBuf))
        go response(ln, addr, requestBuf)
    }
}

func response(server net.PacketConn, addr net.Addr, buf []byte) {
    time := time.Now().Format(time.ANSIC)
    responseStr := fmt.Sprintf("time received: %v. Your message: %v!", time, string(buf))

	server.WriteTo([]byte(responseStr), addr)
}

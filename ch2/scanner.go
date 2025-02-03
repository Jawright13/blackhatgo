package main

import (
	"fmt"
	"net"
)

// Iterative scanning example
func concurrent_scan(address string) {
	for i := 1; i < 1024; i++ {
		addr := fmt.Sprintf("%s:%d", address, i)
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			continue
		}
		conn.Close()
		fmt.Printf("%d open\n", i)
	}
}

func main() {
	address := "scanme.nmap.org"
	concurrent_scan(address)
}

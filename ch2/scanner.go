package main

import (
	"fmt"
	"net"
)

// Iterative scanning example
func non_concurrent_scan(address string) {
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

// Too fast of a concurrent scanner
func concurrent_scan_toofast(address string) {
	for i := 1; i < 1024; i++ {
		go func(j int) {
			addr := fmt.Sprintf("%s:%d", address, i)
			conn, err := net.Dial("tcp", addr)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("%d open\n", i)
		}(i)
	}
}

func main() {
	address := "scanme.nmap.org"
	//non_concurrent_scan(address)
	concurrent_scan_toofast(address)
}

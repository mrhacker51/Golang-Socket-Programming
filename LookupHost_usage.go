package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage : %s ipAddr\n", os.Args[0])
		os.Exit(0)
	}
	address := os.Args[1]

	host, err := net.LookupHost(address)

	if err != nil {
		log.Fatal("Failed ...")
		os.Exit(0)
	}

	for _, v := range host {
		fmt.Fprintf(os.Stderr, "Host : %s\n", v)
	}

	os.Exit(0)
}

/*
RESULT 

Host : 216.58.214.132 // ipv4
Host : 2a00:1450:4017:800::2004 // ipv6 address

*/

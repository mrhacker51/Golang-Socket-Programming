package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage : %s NetworkType - Service\n", os.Args[0])
		os.Exit(0)
	}

	// NetworkType == TCP , UDP
	// Service == Telnet , SSH Vb...
	var (
		networktype = os.Args[1]
		service     = os.Args[2]
	)

	port, err := net.LookupPort(networktype, service)

	if err != nil {
		fmt.Println("Port Not Found  ...")
		os.Exit(0)
	}

	fmt.Println("Port : ", port)

	os.Exit(0)

}

/*

USAGE : go run main.go tcp ssh   example

*/

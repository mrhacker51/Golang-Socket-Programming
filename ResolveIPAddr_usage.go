package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage %s : Ip Addr\n", os.Args[0])
		os.Exit(1)
	}

	address := os.Args[1]

	host, err := net.ResolveIPAddr("ip", address) // "ip" , "ip4", "ip6" kullanılabilir ...
	if err != nil {
		log.Printf("%s", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Host : %s\n", host.String())
	fmt.Printf("Host : %s\n", host)
	fmt.Printf("Host : %s\n", host.Network())
	fmt.Printf("Host : %s\n", host.IP)
}

/*
RESULT

Host : 216.58.212.4
Host : 216.58.212.4
Host : ip
Host : 216.58.212.4
*/

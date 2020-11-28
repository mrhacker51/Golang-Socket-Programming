package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage : %s Address\n", os.Args[0])
		os.Exit(0)
	}

	name := os.Args[1]

	txt, err := net.LookupTXT(name)

	if err != nil {
		log.Fatal("Failed ..", err.Error())
	}

	fmt.Println("===============")
	fmt.Println("[+] LOADING")
	defer fmt.Println("===============")

	for _, v := range txt {
		fmt.Printf("Result : %s\n", v)
	}
}

/*
RESULT = 

===============
[+] LOADING
Result : v=spf1 ip4:185.9.36.116 a mx ptr include:bsilkroad.com ~all
===============

*/

package main

import (
	"flag"
	"fmt"
)

//First we implement flags for cli port scanner

//flags will just be places for user to input info

func main() {
	//CLI Flags
	target := flag.String("target", "", "Target IP or Range with CIDR")
	timeout := flag.Int("timeout", 3, "Timeout in seconds")
	port := flag.Int("port", 80, "Port to scan")
	//workers := flag.Int("pool","2","Number of concurrent workers") Implementation for later
	//scanType := flag.String("scanType","icmp","Type of scan") implementation for later
	flag.Parse()

	if *target == "" {
		//implement err handling after
		fmt.Println("Error:-target flag is required")
		flag.Usage() //Shows help message
		return

	}
	fmt.Println(*target, *timeout, *port)
	return

}

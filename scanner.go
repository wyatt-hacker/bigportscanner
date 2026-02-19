//Multiple concurrent
// scanners that solely access network interface and collect ip
// address saving identification and parsing
// for parser and report builder

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func getArpEntries() (map[string]string, error) {
	// read proc/net/arp file and parse by whitte space segments to grab IP address and MAC Address
	data, err := ioutil.ReadFile("/proc/net/arp")
	if err != nil {
		return nil, fmt.Errorf("unable to read ArpTable: %w", err)

	}
	arpMap := make(map[string]string)
	// Map will have ip and mac address
	tableEntries := strings.Split(string(data), "\n")
	// entries in arp table are split by white space, i need to  grab IP and MAC now
	// loop and collect slices that match IP and MAC
	for i := 1; i < len(tableEntries); i++ {
		entry := tableEntries[i]

		if len(entry) == 0 {
			continue
		}
		//extract IP and MAC and add to MAc
		fields := strings.Fields(entry)

		if len(fields) < 4 {
			fmt.Println("not enough info in ARP table")

		}

		ip := fields[0]
		mac := fields[3]

		arpMap[ip] = mac
	}

	return arpMap, nil

}

func main() {
	result, err := getArpEntries()

	if err != nil {
		fmt.Print("no results found??")
		return
	} else {
		for ip, mac := range result {
			fmt.Printf("IP: %s MAC: %s\n", ip, mac)
		}
	}
}

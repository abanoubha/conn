package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Ullaakut/nmap/v2"
)

func main() {
	ds := runNmap()

	fmt.Printf("connected devices : %d \n", len(ds)-1)

	for i, d := range ds {
		fmt.Println(i, d)
	}
}

func runNmap() []string {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	scanner, err := nmap.NewScanner(
		nmap.WithCustomArguments("-sn", "192.168.1.0/24"),
		nmap.WithContext(ctx),
	)
	if err != nil {
		log.Fatalf("unable to create nmap scanner: %v", err)
	}

	result, warnings, err := scanner.Run()
	if err != nil {
		log.Fatalf("unable to run nmap scan: %v", err)
	}

	if warnings != nil {
		log.Printf("Warnings: \n %v", warnings)
	}

	// Use the results to print an example output
	for _, host := range result.Hosts {
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			continue
		}

		fmt.Printf("Host %q:\n", host.Addresses[0])

		for _, port := range host.Ports {
			fmt.Printf("\tPort %d/%s %s %s\n", port.ID, port.Protocol, port.State, port.Service.Name)
		}
	}

	//fmt.Printf("connected devices : %d \n", len(result.Hosts))

	// for _, host := range result.Hosts {
	// 	fmt.Println(host.Addresses[0])
	// }

	devices := []string{}

	for _, host := range result.Hosts {
		devices = append(devices, host.Addresses[0].Addr)
	}

	return devices
}

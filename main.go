package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"github.com/Ullaakut/nmap/v2"
)

func main() {
	go func() {
		w := app.NewWindow(
			app.Size(unit.Dp(800), unit.Dp(400)),
			app.Title("conn : list all connected devices"),
		)

		ds := runNmap()
		num := len(ds) - 1

		if err := loop(w, num, ds); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func loop(w *app.Window, num int, ds []string) error {
	var ops op.Ops
	for e := range w.Events() {
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err

		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)
			drawTable(gtx, num, ds)
			e.Frame(gtx.Ops)
		}
	}
	return nil
}

func drawTable(gtx layout.Context, num int, ds []string) {
	fmt.Println(strconv.Itoa(num), ds[0])
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

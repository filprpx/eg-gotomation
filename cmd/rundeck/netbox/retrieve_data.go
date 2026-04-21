package main

import (
	"context"
	"fmt"
	"log"

	"github.com/filprpx/eg-gotomation/internal/netbox"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	netbox, err := netbox.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	devices, err := netbox.DCIM.Device.List(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Retrieved %d devices from netbox:\n", len(devices))
	for i, device := range devices {
		fmt.Printf("%d. %s\n", i+1, device.Name)
	}
}

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/filprpx/eg-gotomation/internal/netbox"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
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

		fmt.Printf("retriving data...\n")

		dev_info, err := netbox.DCIM.Device.Get(ctx, device.Id)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("device info retrieved!\n")
		fmt.Printf("device_type for this device is: %s\n", dev_info.DeviceType.Display)

		fmt.Printf("testing update\n")

		device.Description = "teste 1"
		ok, err := netbox.DCIM.Device.Update(ctx, &device)
		if err != nil {
			log.Fatal(err)
		}

		if ok {
			fmt.Println("Device updated successfuly!")
		}
	}
}
